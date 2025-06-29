load("@bazel_skylib//rules:expand_template.bzl", "expand_template")
load("@bazel_skylib//rules:write_file.bzl", "write_file")
load("@rules_cc//cc:defs.bzl", "cc_binary", "cc_library")

write_file(
    name = "config.h_template",
    out = "config.h.in",
    content = ["""\
#define LINUX
#define DEBUG 0
#define USE_PTHREAD_ASYNCDNS

#define CHRONY_VERSION "%CHRONY_VERSION%"
#define CHRONYD_FEATURES "NTP RTC SCFILTER ASYNCDNS"
#define FEAT_NTP
#define FEAT_RTC
#define FEAT_SCFILTER
#define FEAT_ASYNCDNS
#define FEAT_PRIVDROP
#define CAP_IS_SUPPORTED

#define DEFAULT_COMMAND_SOCKET "/todo/chronyd.sock"
#define DEFAULT_CONF_FILE "/todo/chrony.conf"
#define DEFAULT_HWCLOCK_FILE ""
#define DEFAULT_PID_FILE "/todo/chronyd.pid"
#define DEFAULT_RTC_DEVICE "/dev/rtc"
#define DEFAULT_USER "root"

#define MAIL_PROGRAM "/todo/sendmail"
"""],
)

expand_template(
    name = "config.h_expanded",
    template = ":config.h.in",
    out = "config.h",
    substitutions = {
        # ONCHANGE(//build/bazel:third_party.MODULE.bazel): version needs to be kept in sync
        "%CHRONY_VERSION%": "4.1-monogon",
    },
)

# Headers which couldn't be decoupled into sub-libraries.
cc_library(
    name = "common_hdrs",
    srcs = [
        ":config.h_expanded",

        # Headers corresponding to .c files in :common.
        "addrfilt.h",
        "array.h",
        "clientlog.h",
        "cmdparse.h",
        "conf.h",
        "keys.h",
        "local.h",
        "logging.h",
        "memory.h",
        "nameserv.h",
        "reference.h",
        "regress.h",
        "samplefilt.h",
        "sched.h",
        "smooth.h",
        "socket.h",
        "sources.h",
        "sourcestats.h",
        "util.h",

        # Corresponding to .c files in :ntp.
        "ntp_core.h",
        "ntp_sources.h",
        "nts_ke.h",

        # Other headers.
        "addressing.h",
        "candm.h",
        "cmdmon.h",
        "cmac.h",
        "hash.h",
        "localp.h",
        "manual.h",
        "ntp.h",
        "privops.h",
        "refclock.h",
        "reports.h",
        "siv.h",
        "srcparams.h",
        "sysincl.h",
    ],
)

# Sources which couldn't be decoupled into sub-libraries.
cc_library(
    name = "common",
    srcs = [
        "addrfilt.c",
        "array.c",
        "clientlog.c",
        "cmdparse.c",
        "conf.c",
        "keys.c",
        "local.c",
        "logging.c",
        "memory.c",
        "reference.c",
        "regress.c",
        "samplefilt.c",
        "sched.c",
        "smooth.c",
        "socket.c",
        "sources.c",
        "sourcestats.c",
        "util.c",
    ],
    deps = [
        ":common_hdrs",
    ],
)

# MD5 library used by keys.c, which does #include "md5.c".
cc_library(
    name = "md5",
    textual_hdrs = [
        "md5.h",
        "md5.c",
    ],
)

cc_library(
    name = "nameserv",
    srcs = [
        "nameserv.c",
        "nameserv_async.h",
        "nameserv_async.c",
    ],
    deps = [
        ":common",
    ],
)

cc_library(
    name = "ntp",
    srcs = [
        "nts_ke_client.h",
        "nts_ke_server.h",
        "nts_ke_session.h",
        "nts_ntp_client.h",
        "nts_ntp_auth.h",
        "nts_ntp_server.h",
        "nts_ntp.h",
        "ntp_auth.h",
        "ntp_auth.c",
        "ntp_core.c",
        "ntp_ext.h",
        "ntp_ext.c",
        "ntp_io.h",
        "ntp_io.c",
        "ntp_signd.h",
        "ntp_sources.c",
    ],
    deps = [
        ":common",
        ":nameserv",
    ],
)

cc_library(
    name = "sys",
    srcs = [
        "sys.h",
        "sys.c",
        "sys_generic.h",
        "sys_generic.c",
        "sys_linux.h",
        "sys_linux.c",
        "sys_timex.h",
        "sys_timex.c",
        "sys_posix.h",
        "sys_null.h",
        "sys_null.c",
    ],
    deps = [
        ":common",
        "@seccomp//:seccomp",
        "@libcap//:libcap",
    ],
)

cc_library(
    name = "rtc",
    srcs = [
        "rtc.h",
        "rtc.c",
        "rtc_linux.h",
        "rtc_linux.c",
    ],
    deps = [
        ":common",
        ":sys",
    ],
)

cc_library(
    name = "tempcomp",
    srcs = [
        "tempcomp.h",
        "tempcomp.c",
    ],
    deps = [
        ":common",
    ],
)

cc_binary(
    name = "chrony",
    srcs = [
        "hash_intmd5.c",
        "main.h",
        "main.c",
        "stubs.c",
    ],
    deps = [
        ":common",
        ":md5",
        ":ntp",
        ":rtc",
        ":tempcomp",
    ],
    visibility = ["//visibility:public"],
)
