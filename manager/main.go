package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "strconv"
    "net/http"

    "docker.io/go-docker"
    "docker.io/go-docker/api/types"
    "docker.io/go-docker/api/types/swarm"
)

var nextPort = 5001

func pageNew(w http.ResponseWriter, r *http.Request) {
    cli, err := docker.NewEnvClient()
    if err != nil {
        panic(err)
	}

    newPort := nextPort
    nextPort++

    newHost := "http://" + strings.Split(r.Host, ":")[0] + ":" + strconv.Itoa(newPort) + "/"

    _, err = cli.ServiceCreate(context.Background(), swarm.ServiceSpec{
        TaskTemplate: swarm.TaskSpec{
            ContainerSpec: &swarm.ContainerSpec{
                Image: "slanewalsh/swarmchat:worker",
            },
            RestartPolicy: &swarm.RestartPolicy{
                Condition: swarm.RestartPolicyConditionNone,
            },
        },
        EndpointSpec: &swarm.EndpointSpec{
            Mode: swarm.ResolutionModeVIP,
            Ports: []swarm.PortConfig{
                swarm.PortConfig{
                    Name: "http",
                    Protocol: swarm.PortConfigProtocolTCP,
                    TargetPort: 80,
                    PublishedPort: uint32(newPort),
                },
            },
        },
    }, types.ServiceCreateOptions{})

    if err != nil {
        log.Fatal(err)
    }

    //http.Redirect(w, r, newHost, http.StatusFound)
    fmt.Fprintf(w, `
<html>
<head>
    <title></title>
</head>
<body>
    Redirecting...
    <script>
    window.setTimeout(function() {
        location.href = "%s";
    }, 5000);
    </script>
</body>
</html>
    `, newHost)
}

func main() {
    fs := http.FileServer(http.Dir("public"))
    http.Handle("/", fs)

    http.HandleFunc("/new", pageNew)
    log.Fatal(http.ListenAndServe(":80", nil))
}
