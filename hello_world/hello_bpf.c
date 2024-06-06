#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

SEC("kprobe/sys_execve")
int hello(void *ctx) {
    u32 pid = bpf_get_current_pid_tgid() >> 32;
    u32 uid = bpf_get_current_uid_gid();
    if (uid == 0) {
    bpf_printk("Hello World! PID: %u, UID: %u\n", pid, uid);
    }
    return 0;
}

char _license[] SEC("license") = "GPL";
