func getPodLogs(pod corev1.Pod) string {
    podLogOpts := corev1.PodLogOptions{}
    config, err := rest.InClusterConfig()
    if err != nil {
        return "error in getting config"
    }
    // creates the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return "error in getting access to K8S"
    }
    req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)
    podLogs, err := req.Stream()
    if err != nil {
        return "error in opening stream"
    }
    defer podLogs.Close()

    buf := new(bytes.Buffer)
    _, err = io.Copy(buf, podLogs)
    if err != nil {
        return "error in copy information from podLogs to buf"
    }
    str := buf.String()

    return str
}