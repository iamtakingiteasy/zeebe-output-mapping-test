//go:generate go-bindata -pkg main dia.bpmn
//go:generate go fmt ./...

package main

import (
	"github.com/zeebe-io/zeebe/clients/go/entities"
	"github.com/zeebe-io/zeebe/clients/go/pb"
	"github.com/zeebe-io/zeebe/clients/go/worker"
	"github.com/zeebe-io/zeebe/clients/go/zbc"
	"log"
	"os"
    "sync/atomic"
)

var zbclient zbc.ZBClient

func init() {
	var zeebeUrl string
	if url, ok := os.LookupEnv("ZEEBE_URL"); ok {
		zeebeUrl = url
	} else {
		log.Fatalln("ZEEBE_URL env variable required")
	}

	client, err := zbc.NewZBClient(zeebeUrl)
	if err != nil {
		log.Fatalln(err)
	}
	zbclient = client
}

func startProcess(id string) {
	cmd, err := zbclient.NewCreateInstanceCommand().BPMNProcessId("test-process").LatestVersion().VariablesFromMap(map[string]interface{}{
		"id": id,
	})
	if err != nil {
		panic(err)
	}
	_, err = cmd.Send()
	if err != nil {
		panic(err)
	}
}

func main() {
	bs, err := Asset("dia.bpmn")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = zbclient.NewDeployWorkflowCommand().AddResource(bs, "dia.bpmn", pb.WorkflowRequestObject_BPMN).Send()
	if err != nil {
		log.Fatalln(err)
	}

    var counter int32
    exiter := make(chan int)

	prop := zbclient.NewJobWorker().JobType("test-service").Handler(func(client worker.JobClient, job entities.Job) {
        defer func() {
            if atomic.AddInt32(&counter, 1) == 2 {
                log.Println("Done, exiting")
                close(exiter)
            }
        }()
		vars, err := job.GetVariablesAsMap()
		if err != nil {
			log.Fatalln(err)
		}
        log.Printf("worker got: %v\n", vars)
		var id string
		if v, ok := vars["id"].(string); ok {
			id = v
		}
		var payload string
		if id == "123" {
			payload = `{"account":{"id":"123","test1":"","test2":"success"}, "foo": {"unmapped_values_do_not_affect":""}}`
		} else {
			payload = `{"account":{"id":"123","test1":"","test2":""}}`
		}
        log.Printf("worker return: %v\n", payload)
		cmd, err := client.NewCompleteJobCommand().JobKey(job.GetKey()).VariablesFromString(payload)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = cmd.Send()
		if err != nil {
			log.Fatalln(err)
		}
	})
	prop.Name("test-service-1")
	prop.Concurrency(1)
	prop.Open()

	go func() {
        log.Println("running abc")
		startProcess("abc")
        log.Println("running 123")
		startProcess("123")
	}()
    <-exiter
}
