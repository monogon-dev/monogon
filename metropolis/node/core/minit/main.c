// minit is a barebones Linux-compatible init (PID 1) process.
//
// Its goal is to run the Metropolis core executable and reap any children that
// it stumbles upon. It does not support running under a TTY and is not
// configurable in any way.
//
// The only reason this exists is because Go's child process reaping (when
// using os/exec.Command) races any PID 1 process reaping, thereby preventing
// running a complex Go binary as PID 1. In the future this might be rewritten
// in a memory-safe language like Zig or Rust, but this implementation will do
// for now, as long as it keeps having basically zero attack surface.
//
// This code has been vaguely inspired by github.com/Yelp/dumb-init and
// github.com/krallin/tini, two already existing minimal init implementations.
// These, however, attempt to handle being run in a TTY and some
// configurability, as they're meant to be run in containers. We don't need any
// of that, and we'd rather have as little C as possible.

#include <errno.h>
#include <linux/reboot.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/reboot.h>
#include <sys/wait.h>
#include <unistd.h>

void handle_signal(pid_t child_pid, int signum);

int main() {
    // Block all signals. We'll unblock them in the child.
    sigset_t all_signals;
    sigfillset(&all_signals);
    sigprocmask(SIG_BLOCK, &all_signals, NULL);

    // Say hello.
    fprintf(stderr,
        "\n"
        "  Metropolis Cluster Operating System\n"
        "  Copyright 2020-2021 The Monogon Project Authors\n"
        "\n"
    );


    pid_t pid = fork();
    if (pid < 0) {
        fprintf(stderr, "fork(): %s\n", strerror(errno));
        return 1;
    }

    if (pid == 0) {
        // In the child. Unblock all signals.
        sigprocmask(SIG_UNBLOCK, &all_signals, NULL);
        if (setsid() == -1) {
            fprintf(stderr, "setsid: %s\n", strerror(errno));
            return 1;
        }

        // Then, start the core executable.
        char *argv[] = {
            "/core",
            NULL,
        };
        execvp(argv[0], argv);
        fprintf(stderr, "execvpe(/core) failed: %s\n", strerror(errno));
        return 1;
    }

    // In the parent. Wait for any signal, then handle it and any other pending
    // ones.
    for (;;) {
        int signum;
        sigwait(&all_signals, &signum);
        handle_signal(pid, signum);
    }
}

// handle_signal is called by the main reap loop for every signal received. It
// reaps children if SIGCHLD is received, and otherwise dispatches the signal to
// its direct child.
void handle_signal(pid_t child_pid, int signum) {
    // Anything other than SIGCHLD should just be forwarded to the child.
    if (signum != SIGCHLD) {
        kill(-child_pid, signum);
        return;
    }

    // A SIGCHLD was received. Go through all children and reap them, checking
    // if any of them is our direct child.

    // exit_status will be set if the direct child process exited.
    int exit_status = -1;

    pid_t killed_pid;
    int status;
    while ((killed_pid = waitpid(-1, &status, WNOHANG)) > 0) {
        if (killed_pid != child_pid) {
            // Something else than our direct child died, just reap it.
            continue;
        }

        // Our direct child exited. Translate its status into an exit code.
        if (WIFEXITED(status)) {
            // For processes which exited, just use the exit code directly.
            exit_status = WEXITSTATUS(status);
        } else if (WIFSIGNALED(status)) {
            // Otherwise, emulate what sh/bash do and return 128 + the signal
            // number that the child received.
            exit_status = 128 + WTERMSIG(status);
        } else {
            // Something unexpected happened. Attempt to handle this gracefully,
            // but complain.
            fprintf(stderr, "child status not EXITED nor SIGNALED: %d\n", status);
            exit_status = 1;
        }
    }

    // Direct child exited, let's also exit.
    if (exit_status >= 0) {
        fprintf(stderr, "\n  Metropolis core exited with status: %d\n", exit_status);
        sync();
        if (exit_status != 0) {
            fprintf(stderr, "  Disks synced, rebooting in 30 seconds...\n", exit_status);
            sleep(30);
            fprintf(stderr, "  Rebooting...\n\n", exit_status);
        } else {
            fprintf(stderr, "  Disks synced, rebooting...\n\n");
        }
        reboot(LINUX_REBOOT_CMD_RESTART);
    }
}
