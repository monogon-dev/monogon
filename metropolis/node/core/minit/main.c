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
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/reboot.h>
#include <sys/wait.h>
#include <unistd.h>

void handle_signal(pid_t child_pid, int signum);

#define NUM_CONSOLES 3
FILE *consoles[NUM_CONSOLES] = {};

// open_consoles populates the consoles array with FILE pointers to opened
// character devices that should receive log messages. Some of these pointers
// are likely to be null, meaning that particular console is not available.
void open_consoles() {
    consoles[0] = fopen("/dev/console", "w");
    consoles[1] = fopen("/dev/tty0", "w");
    consoles[2] = fopen("/dev/ttyS0", "w");

    // Set all open consoles to be line-buffered.
    for (int i = 0; i < NUM_CONSOLES; i++) {
        if (consoles[i] == NULL) {
            continue;
        }
        setvbuf(consoles[i], NULL, _IOLBF, BUFSIZ);
    }

    // TODO(q3k): disable hardware and software flow control on TTYs. This
    // shouldn't be necessary on our current platform, but should be ensured
    // regardless, to make sure we never block writing to any console.
}

// cprintf emits a format string to all opened consoles.
void cprintf(const char *fmt, ...) {
    va_list args;
    va_start(args, fmt);

    for (int i = 0; i < NUM_CONSOLES; i++) {
        FILE *console = consoles[i];
        if (console == NULL) {
            continue;
        }
        vfprintf(console, fmt, args);
    }

    va_end(args);
}

int main() {
    // Block all signals. We'll unblock them in the child.
    sigset_t all_signals;
    sigfillset(&all_signals);
    sigprocmask(SIG_BLOCK, &all_signals, NULL);

    open_consoles();

    // Say hello.
    cprintf(
        "\n"
        "  Metropolis Cluster Operating System\n"
        "  Copyright 2020-2022 The Monogon Project Authors\n"
        "\n"
    );


    pid_t pid = fork();
    if (pid < 0) {
        cprintf("fork(): %s\n", strerror(errno));
        return 1;
    }

    if (pid == 0) {
        // In the child. Unblock all signals.
        sigprocmask(SIG_UNBLOCK, &all_signals, NULL);
        if (setsid() == -1) {
            cprintf("setsid: %s\n", strerror(errno));
            return 1;
        }

        // Then, start the core executable.
        char *argv[] = {
            "/core",
            NULL,
        };
        execvp(argv[0], argv);
        cprintf("execvpe(/core) failed: %s\n", strerror(errno));
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
            cprintf("child status not EXITED nor SIGNALED: %d\n", status);
            exit_status = 1;
        }
    }

    // Direct child exited, let's also exit.
    if (exit_status >= 0) {
        cprintf("\n  Metropolis core exited with status: %d\n", exit_status);
        sync();
        if (exit_status != 0) {
            cprintf("  Disks synced, rebooting in 30 seconds...\n", exit_status);
            sleep(30);
            cprintf("  Rebooting...\n\n", exit_status);
        } else {
            cprintf("  Disks synced, rebooting...\n\n");
        }
        reboot(LINUX_REBOOT_CMD_RESTART);
    }
}
