/*
ALERT CPUUsageHigh
    IF sum(rate(container_cpu_usage_seconds_total{id="/docker/4ac59be542ef32248b5f91f7d1e84337c420132913bb9a0b0f7c3a2c3ea50418"}[5m])) > 0.450
    FOR 5m
    LABELS { severity = "warning", kubernetes_container_name = "nginx", kubernetes_pod_name = "nginx-pc-535175895-sxhet", kubernetes_namespace = "allen", id = "/docker/4ac59be542ef32248b5f91f7d1e84337c420132913bb9a0b0f7c3a2c3ea50418"}
    ANNOTATIONS {
        summary = "Container allen/nginx-pc-535175895-sxhet/nginx cpu usage high",
        description = "Container allen/nginx-pc-535175895-sxhet/nginx has high cpu usage of {{ $value }}",
    }

ALERT CPUUsageHigh
    IF sum(rate(container_cpu_usage_seconds_total{id="/docker/4ac59be542ef32248b5f91f7d1e84337c420132913bb9a0b0f7c3a2c3ea50418"}[5m])) > 0.475
    FOR 5m
    LABELS { severity = "critical", kubernetes_container_name = "nginx", kubernetes_pod_name = "nginx-pc-535175895-sxhet", kubernetes_namespace = "allen", id = "/docker/4ac59be542ef32248b5f91f7d1e84337c420132913bb9a0b0f7c3a2c3ea50418"}
    ANNOTATIONS {
        summary = "Container allen/nginx-pc-535175895-sxhet/nginx cpu usage high",
        description = "Container allen/nginx-pc-535175895-sxhet/nginx has high cpu usage of {{ $value }}",
    }

ALERT MemoryUsageHigh
    IF sum(container_memory_usage_bytes{id="/docker/4ac59be542ef32248b5f91f7d1e84337c420132913bb9a0b0f7c3a2c3ea50418"}) > 241591904
    FOR 5m
    LABELS { severity = "warning", kubernetes_container_name = "nginx", kubernetes_pod_name = "nginx-pc-535175895-sxhet", kubernetes_namespace = "allen", id = "/docker/4ac59be542ef32248b5f91f7d1e84337c420132913bb9a0b0f7c3a2c3ea50418"}
    ANNOTATIONS {
        summary = "Container allen/nginx-pc-535175895-sxhet/nginx memory usage high",
        description = "Container allen/nginx-pc-535175895-sxhet/nginx has high memory usage of {{ $value }}",
    }

ALERT MemoryUsageHigh
    IF sum(container_memory_usage_bytes{id="/docker/4ac59be542ef32248b5f91f7d1e84337c420132913bb9a0b0f7c3a2c3ea50418"}) > 255013680
    FOR 5m
    LABELS { severity = "critical", kubernetes_container_name = "nginx", kubernetes_pod_name = "nginx-pc-535175895-sxhet", kubernetes_namespace = "allen", id = "/docker/4ac59be542ef32248b5f91f7d1e84337c420132913bb9a0b0f7c3a2c3ea50418"}
    ANNOTATIONS {
        summary = "Container allen/nginx-pc-535175895-sxhet/nginx memory usage high",
        description = "Container allen/nginx-pc-535175895-sxhet/nginx has high memory usage of {{ $value }}",
    }
*/

package main

import (
    "fmt"
    "regexp"
    "io/ioutil"
)

func main() {
    dat, _ := ioutil.ReadFile("./configmap")
    // fmt.Println(string(dat))
    // re := regexp.MustCompile("ALERT [a-zA-Z]+\n\\s+IF.*\n\\s+FOR\\s[0-9]+[a-z]\n\\s+LABELS {.*}\n\\s+ANNOTATIONS {\n\\s+")
    re := regexp.MustCompile("ALERT.*IF.*")
    arr := re.FindAllString(string(dat), -1)
    for i, m := range arr {
        fmt.Println(i, m)
    }
}