package main

import (
    "fmt"
    "os/exec"
    "time"
)

func run(cmd string, args ...string) error {
    c := exec.Command(cmd, args...)
    c.Stdout = nil
    c.Stderr = nil
    return c.Run()
}

func WriteGit() error {
    date := time.Now().Format(time.RFC3339)

    if err := run("git", "add", "-A"); err != nil {
        return err
    }
    if err := run("git", "commit", "-m", date); err != nil {
        return err
    }
    if err := run("git", "push"); err != nil {
        return err
    }
    return nil
}

func main() {
    if err := WriteGit(); err != nil {
        fmt.Println("error:", err)
    }
}

