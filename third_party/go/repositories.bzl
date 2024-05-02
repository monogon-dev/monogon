load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "af_inet_netaddr",
        importpath = "inet.af/netaddr",
        sum = "h1:U2fwK6P2EqmopP/hFLTOAjWTki0qgd4GMJn5X8wOleU=",
        version = "v0.0.0-20220811202034-502d2d690317",
    )
    go_repository(
        name = "ag_pack_amqp",
        importpath = "pack.ag/amqp",
        sum = "h1:cuNDWLUTbKRtEZwhB0WQBXf9pGbm87pUBXQhvcFxBWg=",
        version = "v0.11.2",
    )
    go_repository(
        name = "cat_dario_mergo",
        importpath = "dario.cat/mergo",
        sum = "h1:AGCNq9Evsj31mOgNPcLyXc+4PNABt905YmuqPYYpBWk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "cc_mvdan_interfacer",
        importpath = "mvdan.cc/interfacer",
        sum = "h1:WX1yoOaKQfddO/mLzdV4wptyWgoH/6hwLs7QHTixo0I=",
        version = "v0.0.0-20180901003855-c20040233aed",
    )
    go_repository(
        name = "cc_mvdan_lint",
        importpath = "mvdan.cc/lint",
        sum = "h1:DxJ5nJdkhDlLok9K6qO+5290kphDJbHOQO1DFFFTeBo=",
        version = "v0.0.0-20170908181259-adc824a0674b",
    )
    go_repository(
        name = "co_honnef_go_gotraceui",
        importpath = "honnef.co/go/gotraceui",
        sum = "h1:dmNsfQ9Vl3GwbiVD7Z8d/osC6WtGGrasyrC2suc4ZIQ=",
        version = "v0.2.0",
    )
    go_repository(
        name = "co_honnef_go_tools",
        importpath = "honnef.co/go/tools",
        sum = "h1:6qXr+R5w+ktL5UkwEbPp+fEvfyoMPche6GkOpGHZcLc=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_4d63_gocheckcompilerdirectives",
        importpath = "4d63.com/gocheckcompilerdirectives",
        sum = "h1:AHcMYuw56NPjq/2y615IGg2kYkBdTvOaojYCBcRE7MA=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_4d63_gochecknoinits",
        importpath = "4d63.com/gochecknoinits",
        sum = "h1:bf5qocEKjrY58JO2GwywfLsb1199lIVs7qHkiplwHy0=",
        version = "v0.0.0-20200108094044-eb73b47b9fc4",
    )
    go_repository(
        name = "com_github_99designs_gqlgen",
        importpath = "github.com/99designs/gqlgen",
        sum = "h1:7Qc4Ll3mfN3doAyUWOgtGLcBGu+KDgK48HdkBGLZVFs=",
        version = "v0.16.0",
    )
    go_repository(
        name = "com_github_adalogics_go_fuzz_headers",
        importpath = "github.com/AdaLogics/go-fuzz-headers",
        sum = "h1:bvDV9vkmnHYOMsOr4WLk+Vo07yKIzd94sVoIqshQ4bU=",
        version = "v0.0.0-20230811130428-ced1acdcaa24",
    )
    go_repository(
        name = "com_github_adamkorcz_go_118_fuzz_build",
        importpath = "github.com/AdamKorcz/go-118-fuzz-build",
        sum = "h1:59MxjQVfjXsBpLy+dbd2/ELV5ofnUkUZBvWSC85sheA=",
        version = "v0.0.0-20230306123547-8075edf89bb0",
    )
    go_repository(
        name = "com_github_adrg_xdg",
        importpath = "github.com/adrg/xdg",
        sum = "h1:RzRqFcjH4nE5C6oTAxhBtoE2IRyjBSa62SCbyPidvls=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_afex_hystrix_go",
        importpath = "github.com/afex/hystrix-go",
        sum = "h1:rFw4nCn9iMW+Vajsk51NtYIcwSTkXr+JGrMd36kTDJw=",
        version = "v0.0.0-20180502004556-fa1af6a1f4f5",
    )
    go_repository(
        name = "com_github_agnivade_levenshtein",
        importpath = "github.com/agnivade/levenshtein",
        sum = "h1:n6qGwyHG61v3ABce1rPVZklEYRT8NFpCMrpZdBUbYGM=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_ajstarks_deck",
        importpath = "github.com/ajstarks/deck",
        sum = "h1:7kQgkwGRoLzC9K0oyXdJo7nve/bynv/KwUsxbiTlzAM=",
        version = "v0.0.0-20200831202436-30c9fc6549a9",
    )
    go_repository(
        name = "com_github_ajstarks_deck_generate",
        importpath = "github.com/ajstarks/deck/generate",
        sum = "h1:iXUgAaqDcIUGbRoy2TdeofRG/j1zpGRSEmNK05T+bi8=",
        version = "v0.0.0-20210309230005-c3f852c02e19",
    )
    go_repository(
        name = "com_github_ajstarks_svgo",
        importpath = "github.com/ajstarks/svgo",
        sum = "h1:slYM766cy2nI3BwyRiyQj/Ud48djTMtMebDqepE95rw=",
        version = "v0.0.0-20211024235047-1546f124cd8b",
    )
    go_repository(
        name = "com_github_akavel_rsrc",
        importpath = "github.com/akavel/rsrc",
        sum = "h1:Zxm8V5eI1hW4gGaYsJQUhxpjkENuG91ki8B4zCrvEsw=",
        version = "v0.10.2",
    )
    go_repository(
        name = "com_github_alcortesm_tgz",
        importpath = "github.com/alcortesm/tgz",
        sum = "h1:uSoVVbwJiQipAclBbw+8quDsfcvFjOpI5iCf4p/cqCs=",
        version = "v0.0.0-20161220082320-9c5fe88206d7",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin",
        importpath = "github.com/alecthomas/kingpin",
        sum = "h1:5svnBTFgJjZvGKyYBtMB0+m5wvrbUHiqye8wRJMlnYI=",
        version = "v2.2.6+incompatible",
    )
    go_repository(
        name = "com_github_alecthomas_kingpin_v2",
        importpath = "github.com/alecthomas/kingpin/v2",
        sum = "h1:f48lwail6p8zpO1bC4TxtqACaGqHYA22qkHjHpqDjYY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_github_alecthomas_template",
        importpath = "github.com/alecthomas/template",
        sum = "h1:JYp7IbQjafoB+tBA3gMyHYHrpOtNuDiK/uB5uXxq5wM=",
        version = "v0.0.0-20190718012654-fb15b899a751",
    )
    go_repository(
        name = "com_github_alecthomas_units",
        importpath = "github.com/alecthomas/units",
        sum = "h1:s6gZFSlWYmbqAuRjVTiNNhvNRfY2Wxp9nhfyel4rklc=",
        version = "v0.0.0-20211218093645-b94a6e3cc137",
    )
    go_repository(
        name = "com_github_alexflint_go_filemutex",
        importpath = "github.com/alexflint/go-filemutex",
        sum = "h1:1v0TJPDtlhgpW4nJ+GvxCLSlUDC3+gW0CQQvlmfDR/s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_andybalholm_brotli",
        importpath = "github.com/andybalholm/brotli",
        sum = "h1:8uQZIdzKmjc/iuPu7O2ioW48L81FgatrcpfFmiq/cCs=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_anmitsu_go_shlex",
        importpath = "github.com/anmitsu/go-shlex",
        sum = "h1:kFOfPq6dUM1hTo4JG6LR5AXSUEsOjtdm0kw0FtQtMJA=",
        version = "v0.0.0-20161002113705-648efa622239",
    )
    go_repository(
        name = "com_github_antihax_optional",
        importpath = "github.com/antihax/optional",
        sum = "h1:xK2lYat7ZLaVVcIuj82J8kIro4V6kDe0AUDFboUCwcg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_antlr_antlr4_runtime_go_antlr_v4",
        importpath = "github.com/antlr/antlr4/runtime/Go/antlr/v4",
        sum = "h1:X8MJ0fnN5FPdcGF5Ij2/OW+HgiJrRg3AfHAx1PJtIzM=",
        version = "v4.0.0-20230321174746-8dcc6526cfb1",
    )
    go_repository(
        name = "com_github_antonmedv_expr",
        importpath = "github.com/antonmedv/expr",
        sum = "h1:8YrTtlCzlOtXw+hpeCLDLL2uo0C0k6jmYpYTGws5c5w=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_github_aokoli_goutils",
        importpath = "github.com/aokoli/goutils",
        sum = "h1:7fpzNGoJ3VA8qcrm++XEE1QUe0mIwNeLa02Nwq7RDkg=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_apache_arrow_go_arrow",
        importpath = "github.com/apache/arrow/go/arrow",
        sum = "h1:HGREIyk0QRPt70R69Gm1JFHDgoiyYpCyuGE8E9k/nf0=",
        version = "v0.0.0-20211013220434-5962184e7a30",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v10",
        importpath = "github.com/apache/arrow/go/v10",
        sum = "h1:n9dERvixoC/1JjDmBcs9FPaEryoANa2sCgVFo6ez9cI=",
        version = "v10.0.1",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v11",
        importpath = "github.com/apache/arrow/go/v11",
        sum = "h1:hqauxvFQxww+0mEU/2XHG6LT7eZternCZq+A5Yly2uM=",
        version = "v11.0.0",
    )
    go_repository(
        name = "com_github_apache_arrow_go_v12",
        importpath = "github.com/apache/arrow/go/v12",
        sum = "h1:xtZE63VWl7qLdB0JObIXvvhGjoVNrQ9ciIHG2OK5cmc=",
        version = "v12.0.0",
    )
    go_repository(
        name = "com_github_apache_beam",
        importpath = "github.com/apache/beam",
        sum = "h1:8MOeoZwBgORfaJjrZxpkqJWEIzwupRGLqUqG0/mvEtQ=",
        version = "v2.32.0+incompatible",
    )
    go_repository(
        name = "com_github_apache_thrift",
        importpath = "github.com/apache/thrift",
        sum = "h1:qEy6UW60iVOlUy+b9ZR0d5WzUWYGOo4HfopoyBaNmoY=",
        version = "v0.16.0",
    )
    go_repository(
        name = "com_github_apex_log",
        importpath = "github.com/apex/log",
        sum = "h1:3Zk+boorIQAAGBrHn0JUtAau4ihMamT4WdnfdnXM1zQ=",
        version = "v1.1.4",
    )
    go_repository(
        name = "com_github_apex_logs",
        importpath = "github.com/apex/logs",
        sum = "h1:KmEBVwfDUOTFcBO8cfkJYwdQ5487UZSN+GteOGPmiro=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_aphistic_golf",
        importpath = "github.com/aphistic/golf",
        sum = "h1:2KLQMJ8msqoPHIPDufkxVcoTtcmE5+1sL9950m4R9Pk=",
        version = "v0.0.0-20180712155816-02c07f170c5a",
    )
    go_repository(
        name = "com_github_aphistic_sweet",
        importpath = "github.com/aphistic/sweet",
        sum = "h1:I4z+fAUqvKfvZV/CHi5dV0QuwbmIvYYFDjG0Ss5QpAs=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_apparentlymart_go_cidr",
        importpath = "github.com/apparentlymart/go-cidr",
        sum = "h1:2mAhrMoF+nhXqxTzSZMUzDHkLjmIHC+Zzn4tdgBZjnU=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_armon_circbuf",
        importpath = "github.com/armon/circbuf",
        sum = "h1:QEF07wC0T1rKkctt1RINW/+RMTVmiwxETico2l3gxJA=",
        version = "v0.0.0-20150827004946-bbbad097214e",
    )
    go_repository(
        name = "com_github_armon_consul_api",
        importpath = "github.com/armon/consul-api",
        sum = "h1:G1bPvciwNyF7IUmKXNt9Ak3m6u9DE1rF+RmtIkBpVdA=",
        version = "v0.0.0-20180202201655-eb2c6b5be1b6",
    )
    go_repository(
        name = "com_github_armon_go_metrics",
        importpath = "github.com/armon/go-metrics",
        sum = "h1:B7AQgHi8QSEi4uHu7Sbsga+IJDU+CENgjxoo81vDUqU=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_armon_go_radix",
        importpath = "github.com/armon/go-radix",
        sum = "h1:BUAU3CGlLvorLI26FmByPp2eC2qla6E1Tw+scpcg/to=",
        version = "v0.0.0-20180808171621-7fddfc383310",
    )
    go_repository(
        name = "com_github_armon_go_socks5",
        importpath = "github.com/armon/go-socks5",
        sum = "h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=",
        version = "v0.0.0-20160902184237-e75332964ef5",
    )
    go_repository(
        name = "com_github_aryann_difflib",
        importpath = "github.com/aryann/difflib",
        sum = "h1:pv34s756C4pEXnjgPfGYgdhg/ZdajGhyOvzx8k+23nw=",
        version = "v0.0.0-20170710044230-e206f873d14a",
    )
    go_repository(
        name = "com_github_asaskevich_govalidator",
        importpath = "github.com/asaskevich/govalidator",
        sum = "h1:idn718Q4B6AGu/h5Sxe66HYVdqdGu2l9Iebqhi/AEoA=",
        version = "v0.0.0-20190424111038-f61b66f89f4a",
    )
    go_repository(
        name = "com_github_aws_aws_lambda_go",
        importpath = "github.com/aws/aws-lambda-go",
        sum = "h1:SuCy7H3NLyp+1Mrfp+m80jcbi9KYWAs9/BXwppwRDzY=",
        version = "v1.13.3",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go",
        importpath = "github.com/aws/aws-sdk-go",
        sum = "h1:7JfwifGRGQMHd99PvfXqxBaZsjuRaOF6e3X9zRx2uYo=",
        version = "v1.44.322",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2",
        importpath = "github.com/aws/aws-sdk-go-v2",
        sum = "h1:882kkTpSFhdgYRKVZ/VCgf7sd0ru57p2JCxz4/oN5RY=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_aws_protocol_eventstream",
        importpath = "github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream",
        sum = "h1:dK82zF6kkPeCo8J1e+tGx4JdvDIQzj7ygIoLg8WMuGs=",
        version = "v1.4.10",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_config",
        importpath = "github.com/aws/aws-sdk-go-v2/config",
        sum = "h1:ENTXWKwE8b9YXgQCsruGLhvA9bhg+RqAsL9XEMEsa2c=",
        version = "v1.18.21",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_credentials",
        importpath = "github.com/aws/aws-sdk-go-v2/credentials",
        sum = "h1:oZCEFcrMppP/CNiS8myzv9JgOzq2s0d3v3MXYil/mxQ=",
        version = "v1.13.20",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_ec2_imds",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/ec2/imds",
        sum = "h1:jOzQAesnBFDmz93feqKnsTHsXrlwWORNZMFHMV+WLFU=",
        version = "v1.13.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_feature_s3_manager",
        importpath = "github.com/aws/aws-sdk-go-v2/feature/s3/manager",
        sum = "h1:TnU1cY51027j/MQeFy7DIgk1UuzJY+wLFYqXceY/fiE=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_configsources",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/configsources",
        sum = "h1:dpbVNUjczQ8Ae3QKHbpHBpfvaVkRdesxpTOe9pTouhU=",
        version = "v1.1.32",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_endpoints_v2",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2",
        sum = "h1:QH2kOS3Ht7x+u0gHCh06CXL/h6G8LQJFpZfFBYBNboo=",
        version = "v2.4.26",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_ini",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/ini",
        sum = "h1:gGLG7yKaXG02/jBlg210R7VgQIotiQntNhsCFejawx8=",
        version = "v1.3.34",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_internal_v4a",
        importpath = "github.com/aws/aws-sdk-go-v2/internal/v4a",
        sum = "h1:zsg+5ouVLLbePknVZlUMm1ptwyQLkjjLMWnN+kVs5dA=",
        version = "v1.0.24",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_dynamodb",
        importpath = "github.com/aws/aws-sdk-go-v2/service/dynamodb",
        sum = "h1:0PlAM5X9Tbjr9OpQh3uVIwIbm3kxJpPculFAZQB2u8M=",
        version = "v1.19.4",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_ec2",
        importpath = "github.com/aws/aws-sdk-go-v2/service/ec2",
        sum = "h1:c6a19AjfhEXKlEX63cnlWtSQ4nzENihHZOG0I3wH6BE=",
        version = "v1.93.2",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_eventbridge",
        importpath = "github.com/aws/aws-sdk-go-v2/service/eventbridge",
        sum = "h1:ZRs58K4BH5u8Zzvsy0z9yZlhYW7BsbyUXEsDjy+wZVg=",
        version = "v1.18.9",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_accept_encoding",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding",
        sum = "h1:y2+VQzC6Zh2ojtV2LoC0MNwHWc6qXv/j2vrQtlftkdA=",
        version = "v1.9.11",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_checksum",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/checksum",
        sum = "h1:qIw7Hg5eJEc1uSxg3hRwAthPAO7NeOd4dPxhaTi0yB0=",
        version = "v1.1.27",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_endpoint_discovery",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery",
        sum = "h1:XsLNgECTon/ughUzILFbbeC953tTbXnJv4GQPUHm80A=",
        version = "v1.7.26",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_presigned_url",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/presigned-url",
        sum = "h1:uUt4XctZLhl9wBE1L8lobU3bVN8SNUP7T+olb0bWBO4=",
        version = "v1.9.26",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_internal_s3shared",
        importpath = "github.com/aws/aws-sdk-go-v2/service/internal/s3shared",
        sum = "h1:lRWp3bNu5wy0X3a8GS42JvZFlv++AKsMdzEnoiVJrkg=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_kinesis",
        importpath = "github.com/aws/aws-sdk-go-v2/service/kinesis",
        sum = "h1:bfR+hoEQD1vokNTV1JxSmmaBskT4yI/iF1SjvAYzbvA=",
        version = "v1.17.10",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_s3",
        importpath = "github.com/aws/aws-sdk-go-v2/service/s3",
        sum = "h1:NAc8WQsVQ3+kz3rU619mlz8NcbpZI6FVJHQfH33QK0g=",
        version = "v1.32.0",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sfn",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sfn",
        sum = "h1:u6nKx6nKoDrWVpeLqwMFs2eC4Emn2Fjm+2iZ3+qJQYY=",
        version = "v1.17.9",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sns",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sns",
        sum = "h1:wy1jYAot40/Odzpzeq9S3OfSddJJ5RmpaKujvj5Hz7k=",
        version = "v1.20.8",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sqs",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sqs",
        sum = "h1:SDZBYFUp70hI2T0z9z+KD1iJBz9jGeT7xgU5hPPC9zs=",
        version = "v1.20.8",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sso",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sso",
        sum = "h1:5cb3D6xb006bPTqEfCNaEA6PPEfBXxxy4NNeX/44kGk=",
        version = "v1.12.8",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_ssooidc",
        importpath = "github.com/aws/aws-sdk-go-v2/service/ssooidc",
        sum = "h1:NZaj0ngZMzsubWZbrEFSB4rgSQRbFq38Sd6KBxHuOIU=",
        version = "v1.14.8",
    )
    go_repository(
        name = "com_github_aws_aws_sdk_go_v2_service_sts",
        importpath = "github.com/aws/aws-sdk-go-v2/service/sts",
        sum = "h1:Qf1aWwnsNkyAoqDqmdM3nHwN78XQjec27LjM6b9vyfI=",
        version = "v1.18.9",
    )
    go_repository(
        name = "com_github_aws_smithy_go",
        importpath = "github.com/aws/smithy-go",
        sum = "h1:hgz0X/DX0dGqTYpGALqXJoRKRj5oQ7150i5FdTePzO8=",
        version = "v1.13.5",
    )
    go_repository(
        name = "com_github_aybabtme_rgbterm",
        importpath = "github.com/aybabtme/rgbterm",
        sum = "h1:WWB576BN5zNSZc/M9d/10pqEx5VHNhaQ/yOVAkmj5Yo=",
        version = "v0.0.0-20170906152045-cc83f3b3ce59",
    )
    go_repository(
        name = "com_github_azure_azure_amqp_common_go_v2",
        importpath = "github.com/Azure/azure-amqp-common-go/v2",
        sum = "h1:+QbFgmWCnPzdaRMfsI0Yb6GrRdBj5jVL8N3EXuEUcBQ=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_azure_azure_pipeline_go",
        importpath = "github.com/Azure/azure-pipeline-go",
        sum = "h1:7U9HBg1JFK3jHl5qmo4CTZKFTVgMwdFHMVtCdfBE21U=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_azure_azure_sdk_for_go",
        importpath = "github.com/Azure/azure-sdk-for-go",
        sum = "h1:fcYLmCpyNYRnvJbPerq7U0hS+6+I79yEDJBqVNcqUzU=",
        version = "v68.0.0+incompatible",
    )
    go_repository(
        name = "com_github_azure_azure_service_bus_go",
        importpath = "github.com/Azure/azure-service-bus-go",
        sum = "h1:G1qBLQvHCFDv9pcpgwgFkspzvnGknJRR0PYJ9ytY/JA=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_azure_azure_storage_blob_go",
        importpath = "github.com/Azure/azure-storage-blob-go",
        sum = "h1:1BCg74AmVdYwO3dlKwtFU1V0wU2PZdREkXvAmZJRUlM=",
        version = "v0.14.0",
    )
    go_repository(
        name = "com_github_azure_go_ansiterm",
        importpath = "github.com/Azure/go-ansiterm",
        sum = "h1:UQHMgLO+TxOElx5B5HZ4hJQsoJ/PvUvKRhJHDQXO8P8=",
        version = "v0.0.0-20210617225240-d185dfc1b5a1",
    )
    go_repository(
        name = "com_github_azure_go_autorest",
        importpath = "github.com/Azure/go-autorest",
        sum = "h1:V5VMDjClD3GiElqLWO7mz2MxNAK/vTfRHdAubSIPRgs=",
        version = "v14.2.0+incompatible",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest",
        importpath = "github.com/Azure/go-autorest/autorest",
        sum = "h1:I4+HL/JDvErx2LjyzaVxllw2lRDB5/BT2Bm4g20iqYw=",
        version = "v0.11.29",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_adal",
        importpath = "github.com/Azure/go-autorest/autorest/adal",
        sum = "h1:Yepx8CvFxwNKpH6ja7RZ+sKX+DWYNldbLiALMC3BTz8=",
        version = "v0.9.23",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_azure_auth",
        importpath = "github.com/Azure/go-autorest/autorest/azure/auth",
        sum = "h1:wkAZRgT/pn8HhFyzfe9UnqOjJYqlembgCTi72Bm/xKk=",
        version = "v0.5.12",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_azure_cli",
        importpath = "github.com/Azure/go-autorest/autorest/azure/cli",
        sum = "h1:0W/yGmFdTIT77fvdlGZ0LMISoLHFJ7Tx4U0yeB+uFs4=",
        version = "v0.4.5",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_date",
        importpath = "github.com/Azure/go-autorest/autorest/date",
        sum = "h1:7gUk1U5M/CQbp9WoqinNzJar+8KY+LPI6wiWrP/myHw=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_mocks",
        importpath = "github.com/Azure/go-autorest/autorest/mocks",
        sum = "h1:PGN4EDXnuQbojHbU0UWoNvmu9AGVwYHG9/fkDYhtAfw=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_github_azure_go_autorest_autorest_to",
        importpath = "github.com/Azure/go-autorest/autorest/to",
        sum = "h1:oXVqrxakqqV1UZdSazDOPOLvOIz+XA683u8EctwboHk=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_azure_go_autorest_logger",
        importpath = "github.com/Azure/go-autorest/logger",
        sum = "h1:IG7i4p/mDa2Ce4TRyAO8IHnVhAVF3RFU+ZtXWSmf4Tg=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_azure_go_autorest_tracing",
        importpath = "github.com/Azure/go-autorest/tracing",
        sum = "h1:TYi4+3m5t6K48TGI9AUdb+IzbnSxvnvUMfuitfgcfuo=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_bazelbuild_buildtools",
        importpath = "github.com/bazelbuild/buildtools",
        replace = "github.com/bazelbuild/buildtools",
        sum = "h1:HTepWP/jhtWTC1gvK0RnvKCgjh4gLqiwaOwGozAXcbw=",
        version = "v0.0.0-20231103205921-433ea8554e82",
    )
    go_repository(
        name = "com_github_bazelbuild_rules_go",
        importpath = "github.com/bazelbuild/rules_go",
        sum = "h1:Q+vDhH4yzafZ0xHBT0JEVawb+1nDHUXhjvWTqSGCCyU=",
        version = "v0.43.0",
    )
    go_repository(
        name = "com_github_beevik_ntp",
        importpath = "github.com/beevik/ntp",
        sum = "h1:/w5VhpW5BGKS37vFm1p9oVk/t4HnnkKZAZIubHM6F7Q=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_benbjohnson_clock",
        importpath = "github.com/benbjohnson/clock",
        sum = "h1:VvXlSJBzZpA/zum6Sj74hxwYI2DIxRWuNIoXAzHZz5o=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        sum = "h1:VlbKKnNfV8bJzeqoa4cOKqO6bYr3WgKZxO8Z16+hsOM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bgentry_speakeasy",
        importpath = "github.com/bgentry/speakeasy",
        sum = "h1:ByYyxL9InA1OWqxJqqp2A5pYHUrCiAL6K3J+LKSsQkY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_bitly_go_hostpool",
        importpath = "github.com/bitly/go-hostpool",
        sum = "h1:mXoPYz/Ul5HYEDvkta6I8/rnYM5gSdSV2tJ6XbZuEtY=",
        version = "v0.0.0-20171023180738-a3a6125de932",
    )
    go_repository(
        name = "com_github_bitly_go_simplejson",
        importpath = "github.com/bitly/go-simplejson",
        sum = "h1:6IH+V8/tVMab511d5bn4M7EwGXZf9Hj6i2xSwkNEM+Y=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_bits_and_blooms_bitset",
        importpath = "github.com/bits-and-blooms/bitset",
        sum = "h1:Kn4yilvwNtMACtf1eYDlG8H77R07mZSPbMjLyS07ChA=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_bkaradzic_go_lz4",
        importpath = "github.com/bkaradzic/go-lz4",
        sum = "h1:RXc4wYsyz985CkXXeX04y4VnZFGG8Rd43pRaHsOXAKk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_bketelsen_crypt",
        importpath = "github.com/bketelsen/crypt",
        sum = "h1:+0HFd5KSZ/mm3JmhmrDukiId5iR6w4+BdFtfSy4yWIc=",
        version = "v0.0.3-0.20200106085610-5cbc8cc4026c",
    )
    go_repository(
        name = "com_github_blakesmith_ar",
        importpath = "github.com/blakesmith/ar",
        sum = "h1:m935MPodAbYS46DG4pJSv7WO+VECIWUQ7OJYSoTrMh4=",
        version = "v0.0.0-20190502131153-809d4375e1fb",
    )
    go_repository(
        name = "com_github_blang_semver",
        importpath = "github.com/blang/semver",
        sum = "h1:7hqmJYuaEK3qwVjWubYiht3j93YI0WQBuysxHIfUriU=",
        version = "v3.1.0+incompatible",
    )
    go_repository(
        name = "com_github_blang_semver_v4",
        importpath = "github.com/blang/semver/v4",
        sum = "h1:1PFHFE6yCCTv8C1TeyNNarDzntLi7wMI5i/pzqYIsAM=",
        version = "v4.0.0",
    )
    go_repository(
        name = "com_github_bmizerany_assert",
        importpath = "github.com/bmizerany/assert",
        sum = "h1:DDGfHa7BWjL4YnC6+E63dPcxHo2sUxDIu8g3QgEJdRY=",
        version = "v0.0.0-20160611221934-b7ed37b82869",
    )
    go_repository(
        name = "com_github_boombuler_barcode",
        importpath = "github.com/boombuler/barcode",
        sum = "h1:NDBbPmhS+EqABEs5Kg3n/5ZNjy73Pz7SIV+KCeqyXcs=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_bradfitz_gomemcache",
        importpath = "github.com/bradfitz/gomemcache",
        sum = "h1:hjXJeBcAMS1WGENGqDpzvmgS43oECTx8UXq31UBu0Jw=",
        version = "v0.0.0-20221031212613-62deef7fc822",
    )
    go_repository(
        name = "com_github_bshuster_repo_logrus_logstash_hook",
        importpath = "github.com/bshuster-repo/logrus-logstash-hook",
        sum = "h1:pgAtgj+A31JBVtEHu2uHuEx0n+2ukqUJnS2vVe5pQNA=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_buger_jsonparser",
        importpath = "github.com/buger/jsonparser",
        sum = "h1:2PnMjfWD7wBILjqQbt530v576A/cAbQvEW9gGIpYMUs=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_bugsnag_bugsnag_go",
        importpath = "github.com/bugsnag/bugsnag-go",
        sum = "h1:rFt+Y/IK1aEZkEHchZRSq9OQbsSzIT/OrI8YFFmRIng=",
        version = "v0.0.0-20141110184014-b1d153021fcd",
    )
    go_repository(
        name = "com_github_bugsnag_osext",
        importpath = "github.com/bugsnag/osext",
        sum = "h1:otBG+dV+YK+Soembjv71DPz3uX/V/6MMlSyD9JBQ6kQ=",
        version = "v0.0.0-20130617224835-0dd3f918b21b",
    )
    go_repository(
        name = "com_github_bugsnag_panicwrap",
        importpath = "github.com/bugsnag/panicwrap",
        sum = "h1:nvj0OLI3YqYXer/kZD8Ri1aaunCxIEsOst1BVJswV0o=",
        version = "v0.0.0-20151223152923-e2c28503fcd0",
    )
    go_repository(
        name = "com_github_burntsushi_toml",
        importpath = "github.com/BurntSushi/toml",
        sum = "h1:9F2/+DoOYIOksmaJFPw1tGFy1eDnIJXg+UHjuD8lTak=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_bytecodealliance_wasmtime_go_v14",
        importpath = "github.com/bytecodealliance/wasmtime-go/v14",
        sum = "h1:ur7S3P+PAeJmgllhSrKnGQOAmmtUbLQxb/nw2NZiaEM=",
        version = "v14.0.0",
    )
    go_repository(
        name = "com_github_bytedance_sonic",
        importpath = "github.com/bytedance/sonic",
        sum = "h1:6iJ6NqdoxCDr6mbY8h18oSO+cShGSMRGCEo7F2h0x8s=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_caarlos0_ctrlc",
        importpath = "github.com/caarlos0/ctrlc",
        sum = "h1:2DtF8GSIcajgffDFJzyG15vO+1PuBWOMUdFut7NnXhw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_campoy_unique",
        importpath = "github.com/campoy/unique",
        sum = "h1:V9a67dfYqPLAvzk5hMQOXYJlZ4SLIXgyKIE+ZiHzgGQ=",
        version = "v0.0.0-20180121183637-88950e537e7e",
    )
    go_repository(
        name = "com_github_casbin_casbin_v2",
        importpath = "github.com/casbin/casbin/v2",
        sum = "h1:bTwon/ECRx9dwBy2ewRVr5OiqjeXSGiTUY74sDPQi/g=",
        version = "v2.1.2",
    )
    go_repository(
        name = "com_github_cavaliercoder_go_cpio",
        importpath = "github.com/cavaliercoder/go-cpio",
        sum = "h1:hHg27A0RSSp2Om9lubZpiMgVbvn39bsUmW9U5h0twqc=",
        version = "v0.0.0-20180626203310-925f9528c45e",
    )
    go_repository(
        name = "com_github_cavaliergopher_cpio",
        importpath = "github.com/cavaliergopher/cpio",
        sum = "h1:KQFSeKmZhv0cr+kawA3a0xTQCU4QxXF1vhU7P7av2KM=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_cenkalti_backoff",
        importpath = "github.com/cenkalti/backoff",
        sum = "h1:tNowT99t7UNflLxfYYSlKYsBpXdEet03Pg2g16Swow4=",
        version = "v2.2.1+incompatible",
    )
    go_repository(
        name = "com_github_cenkalti_backoff_v4",
        importpath = "github.com/cenkalti/backoff/v4",
        sum = "h1:MyRJ/UdXutAwSAT+s3wNd7MfTIcy71VQueUuFK343L8=",
        version = "v4.3.0",
    )
    go_repository(
        name = "com_github_census_instrumentation_opencensus_proto",
        importpath = "github.com/census-instrumentation/opencensus-proto",
        sum = "h1:iKLQ0xPNFxR/2hzXZMrBo8f1j86j5WHzznCCQxV/b8g=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_certifi_gocertifi",
        importpath = "github.com/certifi/gocertifi",
        sum = "h1:uH66TXeswKn5PW5zdZ39xEwfS9an067BirqA+P4QaLI=",
        version = "v0.0.0-20200922220541-2c3bb06c6054",
    )
    go_repository(
        name = "com_github_cespare_xxhash",
        importpath = "github.com/cespare/xxhash",
        sum = "h1:a6HrQnmkObjyL+Gs60czilIUGqrzKutQD6XZog3p+ko=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cespare_xxhash_v2",
        importpath = "github.com/cespare/xxhash/v2",
        sum = "h1:DC2CZ1Ep5Y4k3ZQ899DldepgrayRUGE6BBZ/cd9Cj44=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_chai2010_gettext_go",
        importpath = "github.com/chai2010/gettext-go",
        sum = "h1:1Lwwip6Q2QGsAdl/ZKPCwTe9fe0CjlUbqj5bFNSjIRk=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu_v4",
        importpath = "github.com/checkpoint-restore/go-criu/v4",
        sum = "h1:WW2B2uxx9KWF6bGlHqhm8Okiafwwx7Y2kcpn8lCpjgo=",
        version = "v4.1.0",
    )
    go_repository(
        name = "com_github_checkpoint_restore_go_criu_v5",
        importpath = "github.com/checkpoint-restore/go-criu/v5",
        sum = "h1:wpFFOoomK3389ue2lAb0Boag6XPht5QYpipxmSNL4d8=",
        version = "v5.3.0",
    )
    go_repository(
        name = "com_github_chenzhuoyu_base64x",
        importpath = "github.com/chenzhuoyu/base64x",
        sum = "h1:qSGYFH7+jGhDF8vLC+iwCD4WpbV1EBDSzWkJODFLams=",
        version = "v0.0.0-20221115062448-fe3a3abad311",
    )
    go_repository(
        name = "com_github_chzyer_logex",
        importpath = "github.com/chzyer/logex",
        sum = "h1:Swpa1K6QvQznwJRcfTfQJmTE72DqScAa40E+fbHEXEE=",
        version = "v1.1.10",
    )
    go_repository(
        name = "com_github_chzyer_readline",
        importpath = "github.com/chzyer/readline",
        sum = "h1:upd/6fQk4src78LMRzh5vItIt361/o4uq553V8B5sGI=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_chzyer_test",
        importpath = "github.com/chzyer/test",
        sum = "h1:q763qf9huN11kDQavWsoZXJNW3xEE4JJyHa5Q25/sd8=",
        version = "v0.0.0-20180213035817-a1ea475d72b1",
    )
    go_repository(
        name = "com_github_cilium_ebpf",
        importpath = "github.com/cilium/ebpf",
        replace = "github.com/cilium/ebpf",
        sum = "h1:8ht6F9MquybnY97at+VDZb3eQQr8ev79RueWeVaEcG4=",
        version = "v0.12.3",
    )
    go_repository(
        name = "com_github_clbanning_x2j",
        importpath = "github.com/clbanning/x2j",
        sum = "h1:EdRZT3IeKQmfCSrgo8SZ8V3MEnskuJP0wCYNpe+aiXo=",
        version = "v0.0.0-20191024224557-825249438eec",
    )
    go_repository(
        name = "com_github_clickhouse_clickhouse_go",
        importpath = "github.com/ClickHouse/clickhouse-go",
        sum = "h1:iAFMa2UrQdR5bHJ2/yaSLffZkxpcOYQMCUuKeNXGdqc=",
        version = "v1.4.3",
    )
    go_repository(
        name = "com_github_client9_misspell",
        importpath = "github.com/client9/misspell",
        sum = "h1:ta993UF76GwbvJcIo3Y68y/M3WxlpEHPWIGDkJYwzJI=",
        version = "v0.3.4",
    )
    go_repository(
        name = "com_github_cloudflare_golz4",
        importpath = "github.com/cloudflare/golz4",
        sum = "h1:F1EaeKL/ta07PY/k9Os/UFtwERei2/XzGemhpGnBKNg=",
        version = "v0.0.0-20150217214814-ef862a3cdc58",
    )
    go_repository(
        name = "com_github_cncf_udpa_go",
        importpath = "github.com/cncf/udpa/go",
        sum = "h1:QQ3GSy+MqSHxm/d8nCtnAiZdYFd45cYZPs8vOOIYKfk=",
        version = "v0.0.0-20220112060539-c52dc94e7fbe",
    )
    go_repository(
        name = "com_github_cncf_xds_go",
        importpath = "github.com/cncf/xds/go",
        sum = "h1:jQCWAUqqlij9Pgj2i/PB79y4KOPYVyFYdROxgaCwdTQ=",
        version = "v0.0.0-20231128003011-0fa0005c9caa",
    )
    go_repository(
        name = "com_github_cockroachdb_apd",
        importpath = "github.com/cockroachdb/apd",
        sum = "h1:3LFP3629v+1aKXU5Q37mxmRxX/pIu1nijXydLShEq5I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_cockroachdb_cockroach_go_v2",
        importpath = "github.com/cockroachdb/cockroach-go/v2",
        sum = "h1:O7Hl8m0rs/oJNBmRr14ED3Q3+AmugMK9DtJwRDHZ2DA=",
        version = "v2.2.10",
    )
    go_repository(
        name = "com_github_cockroachdb_datadriven",
        importpath = "github.com/cockroachdb/datadriven",
        sum = "h1:H9MtNqVoVhvd9nCBwOyDjUEdZCREqbIdCJD93PBm/jA=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_cockroachdb_errors",
        importpath = "github.com/cockroachdb/errors",
        sum = "h1:Lap807SXTH5tri2TivECb/4abUkMZC9zRoLarvcKDqs=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_github_cockroachdb_logtags",
        importpath = "github.com/cockroachdb/logtags",
        sum = "h1:o/kfcElHqOiXqcou5a3rIlMc7oJbMQkeLk0VQJ7zgqY=",
        version = "v0.0.0-20190617123548-eb05cc24525f",
    )
    go_repository(
        name = "com_github_codahale_hdrhistogram",
        importpath = "github.com/codahale/hdrhistogram",
        sum = "h1:qMd81Ts1T2OTKmB4acZcyKaMtRnY5Y44NuXGX2GFJ1w=",
        version = "v0.0.0-20161010025455-3a0bb77429bd",
    )
    go_repository(
        name = "com_github_codahale_rfc6979",
        importpath = "github.com/codahale/rfc6979",
        sum = "h1:EDmT6Q9Zs+SbUoc7Ik9EfrFqcylYqgPZ9ANSbTAntnE=",
        version = "v0.0.0-20141003034818-6a90f24967eb",
    )
    go_repository(
        name = "com_github_confluentinc_confluent_kafka_go",
        importpath = "github.com/confluentinc/confluent-kafka-go",
        sum = "h1:GCEMecax8zLZsCVn1cea7Y1uR/lRCdCDednpkc0NLsY=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_confluentinc_confluent_kafka_go_v2",
        importpath = "github.com/confluentinc/confluent-kafka-go/v2",
        sum = "h1:qwZtgyGS4OjvebR4TkZPxHAQRN/IbdaxpCQyhDpxeaE=",
        version = "v2.1.1",
    )
    go_repository(
        name = "com_github_container_storage_interface_spec",
        importpath = "github.com/container-storage-interface/spec",
        sum = "h1:D0vhF3PLIZwlwZEf2eNbpujGCNwspwTYf2idJRJx4xI=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_containerd_aufs",
        importpath = "github.com/containerd/aufs",
        sum = "h1:2oeJiwX5HstO7shSrPZjrohJZLzK36wvpdmzDRkL/LY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_btrfs",
        importpath = "github.com/containerd/btrfs",
        sum = "h1:osn1exbzdub9L5SouXO5swW4ea/xVdJZ3wokxN5GrnA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_btrfs_v2",
        importpath = "github.com/containerd/btrfs/v2",
        sum = "h1:FN4wsx7KQrYoLXN7uLP0vBV4oVWHOIKDRQ1G2Z0oL5M=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_containerd_cgroups",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/cgroups",
        sum = "h1:v8rEWFl6EoqHB+swVNjVoCJE8o3jX7e8nqBGPLaDFBM=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_containerd_cgroups_v3",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/cgroups/v3",
        sum = "h1:f5WFqIVSgo5IZmtTT3qVBo6TzI1ON6sycSBKkymb9L0=",
        version = "v3.0.2",
    )
    go_repository(
        name = "com_github_containerd_console",
        importpath = "github.com/containerd/console",
        sum = "h1:F2g4+oChYvBTsASRTz8NP6iIAi97J3TtSAsLbIFn4ro=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_containerd_containerd",
        build_file_proto_mode = "disable",
        build_tags = ["no_btrfs"],  # keep
        importpath = "github.com/containerd/containerd",
        patch_args = ["-p1"],
        pre_patches = [
            # We build with newer CRI API definitions.
            "//third_party/go/patches:containerd-grpc-inherit-unimplemented.patch",
        ],
        sum = "h1:afEHXdil9iAm03BmhjzKyXnnEBtjaLJefdU7DV0IFes=",
        version = "v1.7.15",
    )
    go_repository(
        name = "com_github_containerd_continuity",
        importpath = "github.com/containerd/continuity",
        sum = "h1:v3y/4Yz5jwnvqPKJJ+7Wf93fyWoCB3F5EclWG023MDM=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_github_containerd_fifo",
        importpath = "github.com/containerd/fifo",
        sum = "h1:4I2mbh5stb1u6ycIABlBw9zgtlK8viPI9QkQNRQEEmY=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_containerd_go_cni",
        importpath = "github.com/containerd/go-cni",
        sum = "h1:ORi7P1dYzCwVM6XPN4n3CbkuOx/NZ2DOqy+SHRdo9rU=",
        version = "v1.1.9",
    )
    go_repository(
        name = "com_github_containerd_go_runc",
        importpath = "github.com/containerd/go-runc",
        sum = "h1:oU+lLv1ULm5taqgV/CJivypVODI4SUz1znWjv3nNYS0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_containerd_imgcrypt",
        importpath = "github.com/containerd/imgcrypt",
        sum = "h1:WSf9o9EQ0KGHiUx2ESFZ+PKf4nxK9BcvV/nJDX8RkB4=",
        version = "v1.1.7",
    )
    go_repository(
        name = "com_github_containerd_log",
        importpath = "github.com/containerd/log",
        sum = "h1:TCJt7ioM2cr/tfR8GPbGf9/VRAX8D2B4PjzCpfX540I=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_containerd_nri",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/nri",
        sum = "h1:hdztxwL0gCS1CrCa9bvD1SoJiFN4jBuRQhplCvCPMj8=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_containerd_stargz_snapshotter_estargz",
        importpath = "github.com/containerd/stargz-snapshotter/estargz",
        sum = "h1:OqlDCK3ZVUO6C3B/5FSkDwbkEETK84kQgEeFwDC+62k=",
        version = "v0.14.3",
    )
    go_repository(
        name = "com_github_containerd_ttrpc",
        build_file_proto_mode = "disable",
        importpath = "github.com/containerd/ttrpc",
        sum = "h1:4jlhbXIGvijRtNC8F/5CpuJZ7yKOBFGFOOXg1bkISz0=",
        version = "v1.2.3",
    )
    go_repository(
        name = "com_github_containerd_typeurl",
        importpath = "github.com/containerd/typeurl",
        sum = "h1:Chlt8zIieDbzQFzXzAeBEF92KhExuE4p9p92/QmY7aY=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_containerd_typeurl_v2",
        importpath = "github.com/containerd/typeurl/v2",
        sum = "h1:3Q4Pt7i8nYwy2KmQWIw2+1hTvwTE/6w9FqcttATPO/4=",
        version = "v2.1.1",
    )
    go_repository(
        name = "com_github_containerd_zfs",
        importpath = "github.com/containerd/zfs",
        sum = "h1:n7OZ7jZumLIqNJqXrEc/paBM840mORnmGdJDmAmJZHM=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_containernetworking_cni",
        importpath = "github.com/containernetworking/cni",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:cni-fix-cachepath.patch",
        ],
        sum = "h1:wtRGZVv7olUHMOqouPpn3cXJWpJgM6+EUl31EQbXALQ=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_containernetworking_plugins",
        importpath = "github.com/containernetworking/plugins",
        sum = "h1:SWgg3dQG1yzUo4d9iD8cwSVh1VqI+bP7mkPDoSfP9VU=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_containers_ocicrypt",
        importpath = "github.com/containers/ocicrypt",
        sum = "h1:uoG52u2e91RE4UqmBICZY8dNshgfvkdl3BW6jnxiFaI=",
        version = "v1.1.6",
    )
    go_repository(
        name = "com_github_coredns_caddy",
        importpath = "github.com/coredns/caddy",
        sum = "h1:2eYKZT7i6yxIfGP3qLJoJ7HAsDJqYB+X68g4NYjSrE0=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_coredns_coredns",
        importpath = "github.com/coredns/coredns",
        patch_args = ["-p1"],
        pre_patches = [
            "//third_party/go/patches:coredns-remove-unused-plugins.patch",
        ],
        sum = "h1:IYBM+j/Xx3nTV4HE1s626G9msmJZSdKL9k0ZagYcZFQ=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_github_coredns_corefile_migration",
        importpath = "github.com/coredns/corefile-migration",
        sum = "h1:W/DCETrHDiFo0Wj03EyMkaQ9fwsmSgqTCQDHpceaSsE=",
        version = "v1.0.21",
    )
    go_repository(
        name = "com_github_coreos_bbolt",
        importpath = "github.com/coreos/bbolt",
        sum = "h1:wZwiHHUieZCquLkDL0B8UhzreNWsPHooDAG3q34zk0s=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_coreos_etcd",
        importpath = "github.com/coreos/etcd",
        sum = "h1:8F3hqu9fGYLBifCmRCJsicFqDx/D68Rt3q1JMazcgBQ=",
        version = "v3.3.13+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_etcd",
        importpath = "github.com/coreos/go-etcd",
        sum = "h1:bXhRBIXoTm9BYHS3gE0TtQuyNZyeEMux2sDi4oo5YOo=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_iptables",
        importpath = "github.com/coreos/go-iptables",
        sum = "h1:is9qnZMPYjLd8LYqmm/qlE+wwEgJIkTYdhV3rfZo4jk=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_coreos_go_oidc",
        importpath = "github.com/coreos/go-oidc",
        sum = "h1:mh48q/BqXqgjVHpy2ZY7WnWAbenxRjsz9N1i1YxjHAk=",
        version = "v2.2.1+incompatible",
    )
    go_repository(
        name = "com_github_coreos_go_semver",
        importpath = "github.com/coreos/go-semver",
        sum = "h1:yi21YpKnrx1gt5R+la8n5WgS0kCrsPp33dmEyHReZr4=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_coreos_go_systemd",
        importpath = "github.com/coreos/go-systemd",
        sum = "h1:JOrtw2xFKzlg+cbHpyrpLDmnN1HqhBfnX7WDiW7eG2c=",
        version = "v0.0.0-20190719114852-fd7a80b32e1f",
    )
    go_repository(
        name = "com_github_coreos_go_systemd_v22",
        importpath = "github.com/coreos/go-systemd/v22",
        sum = "h1:RrqgGjYQKalulkV8NGVIfkXQf6YYmOyiJKk8iXXhfZs=",
        version = "v22.5.0",
    )
    go_repository(
        name = "com_github_coreos_pkg",
        importpath = "github.com/coreos/pkg",
        sum = "h1:lBNOc5arjvs8E5mO2tbpBpLoyyu8B6e44T7hJy6potg=",
        version = "v0.0.0-20180928190104-399ea9e2e55f",
    )
    go_repository(
        name = "com_github_corverroos_commentwrap",
        importpath = "github.com/corverroos/commentwrap",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:commentwrap-uri.patch",
        ],
        sum = "h1:toeMwwechJKH0iwOoGJLZK6x42Ba9si+816KxqmgFc8=",
        version = "v0.0.0-20191204065359-2926638be44c",
    )
    go_repository(
        name = "com_github_cosiner_argv",
        importpath = "github.com/cosiner/argv",
        sum = "h1:BVDiEL32lwHukgJKP87btEPenzrrHUjajs/8yzaqcXg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man",
        importpath = "github.com/cpuguy83/go-md2man",
        sum = "h1:BSKMNlYxDvnunlTymqtgONjNnaRV1sTpcovwwjF22jk=",
        version = "v1.0.10",
    )
    go_repository(
        name = "com_github_cpuguy83_go_md2man_v2",
        importpath = "github.com/cpuguy83/go-md2man/v2",
        sum = "h1:qMCsGGgs+MAzDFyp9LpAe1Lqy/fY/qCovCm0qnXZOBM=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_creack_pty",
        importpath = "github.com/creack/pty",
        sum = "h1:n56/Zwd5o6whRC5PMGretI4IdRLlmBXYNjScPaBgsbY=",
        version = "v1.1.18",
    )
    go_repository(
        name = "com_github_crillab_gophersat",
        importpath = "github.com/crillab/gophersat",
        sum = "h1:l4fgnEMmy1+b7pn3nvPwj1ja3Z9MgXE4hUIl9TU8v+M=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_cubicdaiya_gonp",
        importpath = "github.com/cubicdaiya/gonp",
        sum = "h1:ky2uIAJh81WiLcGKBVD5R7KsM/36W6IqqTy6Bo6rGws=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_cyphar_filepath_securejoin",
        importpath = "github.com/cyphar/filepath-securejoin",
        sum = "h1:Ugdm7cg7i6ZK6x3xDF1oEu1nfkyfH53EtKeQYTC3kyg=",
        version = "v0.2.4",
    )
    go_repository(
        name = "com_github_cznic_mathutil",
        importpath = "github.com/cznic/mathutil",
        sum = "h1:iwZdTE0PVqJCos1vaoKsclOGD3ADKpshg3SRtYBbwso=",
        version = "v0.0.0-20181122101859-297441e03548",
    )
    go_repository(
        name = "com_github_cznic_sortutil",
        importpath = "github.com/cznic/sortutil",
        sum = "h1:LpMLYGyy67BoAFGda1NeOBQwqlv7nUXpm+rIVHGxZZ4=",
        version = "v0.0.0-20181122101858-f5f958428db8",
    )
    go_repository(
        name = "com_github_cznic_strutil",
        importpath = "github.com/cznic/strutil",
        sum = "h1:MZRmHqDBd0vxNwenEbKSQqRVT24d3C05ft8kduSwlqM=",
        version = "v0.0.0-20181122101858-275e90344537",
    )
    go_repository(
        name = "com_github_d2g_dhcp4",
        importpath = "github.com/d2g/dhcp4",
        sum = "h1:Xo2rK1pzOm0jO6abTPIQwbAmqBIOj132otexc1mmzFc=",
        version = "v0.0.0-20170904100407-a1d1b6c41b1c",
    )
    go_repository(
        name = "com_github_d2g_dhcp4client",
        importpath = "github.com/d2g/dhcp4client",
        sum = "h1:suYBsYZIkSlUMEz4TAYCczKf62IA2UWC+O8+KtdOhCo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_d2g_dhcp4server",
        importpath = "github.com/d2g/dhcp4server",
        sum = "h1:+CpLbZIeUn94m02LdEKPcgErLJ347NUwxPKs5u8ieiY=",
        version = "v0.0.0-20181031114812-7d4a0a7f59a5",
    )
    go_repository(
        name = "com_github_d2g_hardwareaddr",
        importpath = "github.com/d2g/hardwareaddr",
        sum = "h1:itqmmf1PFpC4n5JW+j4BU7X4MTfVurhYRTjODoPb2Y8=",
        version = "v0.0.0-20190221164911-e7d9fbe030e4",
    )
    go_repository(
        name = "com_github_datadog_appsec_internal_go",
        importpath = "github.com/DataDog/appsec-internal-go",
        sum = "h1:2u5IkF4DBj3KVeQn5Vg2vjPUtt513zxEYglcqnd500U=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_datadog_datadog_agent_pkg_obfuscate",
        importpath = "github.com/DataDog/datadog-agent/pkg/obfuscate",
        sum = "h1:XyYvstMFpSyZtfJHWJm1Sf1meNyCdfhKJrjB6+rUNOk=",
        version = "v0.45.0-rc.1",
    )
    go_repository(
        name = "com_github_datadog_datadog_agent_pkg_remoteconfig_state",
        importpath = "github.com/DataDog/datadog-agent/pkg/remoteconfig/state",
        sum = "h1:PbzrhHhs2+RRdKKti7JBSM8ATIeiji2T2cVt/d8GT8k=",
        version = "v0.48.0-devel.0.20230725154044-2549ba9058df",
    )
    go_repository(
        name = "com_github_datadog_datadog_go_v5",
        importpath = "github.com/DataDog/datadog-go/v5",
        sum = "h1:JLZ6s2K1pG2h9GkvEvMdEGqMDyVLEAccdX5TltWcLMU=",
        version = "v5.1.1",
    )
    go_repository(
        name = "com_github_datadog_go_libddwaf",
        importpath = "github.com/DataDog/go-libddwaf",
        sum = "h1:JgHc+ARmfIzVqEl31HLedVYiNCu3LAQiluvpeNnEx2o=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_datadog_go_tuf",
        importpath = "github.com/DataDog/go-tuf",
        sum = "h1:gld/e3MXfFVB/O8hc3mloP1ayFk75Mmdkmll/9lyd9I=",
        version = "v1.0.1-0.5.2",
    )
    go_repository(
        name = "com_github_datadog_gostackparse",
        importpath = "github.com/DataDog/gostackparse",
        sum = "h1:jb72P6GFHPHz2W0onsN51cS3FkaMDcjb0QzgxxA4gDk=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_datadog_sketches_go",
        importpath = "github.com/DataDog/sketches-go",
        sum = "h1:qTBzWLnZ3kM2kw39ymh6rMcnN+5VULwFs++lEYUUsro=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_datadog_zstd",
        importpath = "github.com/DataDog/zstd",
        sum = "h1:DtpNbljikUepEPD16hD4LvIcmhnhdLTiW/5pHgbmp14=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_github_davecgh_go_spew",
        importpath = "github.com/davecgh/go-spew",
        sum = "h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_daviddengcn_go_colortext",
        importpath = "github.com/daviddengcn/go-colortext",
        sum = "h1:ANqDyC0ys6qCSvuEK7l3g5RaehL/Xck9EX8ATG8oKsE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_decred_dcrd_dcrec_secp256k1_v4",
        importpath = "github.com/decred/dcrd/dcrec/secp256k1/v4",
        sum = "h1:1iy2qD6JEhHKKhUOA9IWs7mjco7lnw2qx8FsRI2wirE=",
        version = "v4.0.0-20210816181553-5444fa50b93d",
    )
    go_repository(
        name = "com_github_denisenkom_go_mssqldb",
        importpath = "github.com/denisenkom/go-mssqldb",
        sum = "h1:9rHa233rhdOyrz2GcP9NM+gi2psgJZ4GWDpL/7ND8HI=",
        version = "v0.11.0",
    )
    go_repository(
        name = "com_github_dennwc_btrfs",
        importpath = "github.com/dennwc/btrfs",
        sum = "h1:ue4Es4Xzz255hWQ7NAWzZxuXG+YOV7URzzusLLSe0zU=",
        version = "v0.0.0-20230312211831-a1f570bd01a1",
    )
    go_repository(
        name = "com_github_dennwc_ioctl",
        importpath = "github.com/dennwc/ioctl",
        sum = "h1:DsWAAjIxRqNcLn9x6mwfuf2pet3iB7aK90K4tF16rLg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_denverdino_aliyungo",
        importpath = "github.com/denverdino/aliyungo",
        sum = "h1:p6poVbjHDkKa+wtC8frBMwQtT3BmqGYBjzMwJ63tuR4=",
        version = "v0.0.0-20190125010748-a747050bb1ba",
    )
    go_repository(
        name = "com_github_derekparker_trie",
        importpath = "github.com/derekparker/trie",
        sum = "h1:G765iDCq7bP5opdrPkXk+4V3yfkgV9iGFuheWZ/X/zY=",
        version = "v0.0.0-20200317170641-1fdf38b7b0e9",
    )
    go_repository(
        name = "com_github_desertbit_timer",
        importpath = "github.com/desertbit/timer",
        sum = "h1:U5y3Y5UE0w7amNe7Z5G/twsBW0KEalRQXZzf8ufSh9I=",
        version = "v0.0.0-20180107155436-c41aec40b27f",
    )
    go_repository(
        name = "com_github_devigned_tab",
        importpath = "github.com/devigned/tab",
        sum = "h1:3mD6Kb1mUOYeLpJvTVSDwSg5ZsfSxfvxGRTxRsJsITA=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_dgrijalva_jwt_go",
        importpath = "github.com/dgrijalva/jwt-go",
        sum = "h1:7qlOGliEKZXTDg6OTjfoBKDXWrumCAMpl/TFQ4/5kLM=",
        version = "v3.2.0+incompatible",
    )
    go_repository(
        name = "com_github_dgryski_go_farm",
        importpath = "github.com/dgryski/go-farm",
        sum = "h1:tdlZCpZ/P9DhczCTSixgIKmwPv6+wP5DGjqLYw5SUiA=",
        version = "v0.0.0-20190423205320-6a90982ecee2",
    )
    go_repository(
        name = "com_github_dgryski_go_rendezvous",
        importpath = "github.com/dgryski/go-rendezvous",
        sum = "h1:lO4WD4F/rVNCu3HqELle0jiPLLBs70cWOduZpkS1E78=",
        version = "v0.0.0-20200823014737-9f7001d12a5f",
    )
    go_repository(
        name = "com_github_dgryski_go_sip13",
        importpath = "github.com/dgryski/go-sip13",
        sum = "h1:RMLoZVzv4GliuWafOuPuQDKSm1SJph7uCRnnS61JAn4=",
        version = "v0.0.0-20181026042036-e10d5fee7954",
    )
    go_repository(
        name = "com_github_dhui_dktest",
        importpath = "github.com/dhui/dktest",
        sum = "h1:0frpeeoM9pHouHjhLeZDuDTJ0PqjDTrycaHaMmkJAo8=",
        version = "v0.3.10",
    )
    go_repository(
        name = "com_github_dimchansky_utfbom",
        importpath = "github.com/dimchansky/utfbom",
        sum = "h1:vV6w1AhK4VMnhBno/TPVCoK9U/LP0PkLCS9tbxHdi/U=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_dimfeld_httptreemux_v5",
        importpath = "github.com/dimfeld/httptreemux/v5",
        sum = "h1:p8jkiMrCuZ0CmhwYLcbNbl7DDo21fozhKHQ2PccwOFQ=",
        version = "v5.5.0",
    )
    go_repository(
        name = "com_github_diskfs_go_diskfs",
        importpath = "github.com/diskfs/go-diskfs",
        sum = "h1:Ow4xorEDw1VNYKbC+SA/qQNwi5gWIwdKUxmUcLFST24=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_distribution_reference",
        importpath = "github.com/distribution/reference",
        sum = "h1:/FUIFXtfc/x2gpa5/VGfiGLuOIdYa1t65IKK2OFGvA0=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_dnaeon_go_vcr",
        importpath = "github.com/dnaeon/go-vcr",
        sum = "h1:zHCHvJYTMh1N7xnV7zf1m1GPBF9Ad0Jk/whtQ1663qI=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_dnstap_golang_dnstap",
        importpath = "github.com/dnstap/golang-dnstap",
        sum = "h1:KRHBoURygdGtBjDI2w4HifJfMAhhOqDuktAokaSa234=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_docker_cli",
        importpath = "github.com/docker/cli",
        sum = "h1:Zcse1DuDqBdgI7OQDV8Go7b83xLgfhW1eza4HfEdxpY=",
        version = "v23.0.3+incompatible",
    )
    go_repository(
        name = "com_github_docker_distribution",
        importpath = "github.com/docker/distribution",
        sum = "h1:T3de5rq0dB1j30rp0sA2rER+m322EBzniBPB6ZIzuh8=",
        version = "v2.8.2+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker",
        importpath = "github.com/docker/docker",
        sum = "h1:9GhVsShNWz1hO//9BNg/dpMnZW25KydO4wtVxWAIbho=",
        version = "v23.0.3+incompatible",
    )
    go_repository(
        name = "com_github_docker_docker_credential_helpers",
        importpath = "github.com/docker/docker-credential-helpers",
        sum = "h1:xtCHsjxogADNZcdv1pKUHXryefjlVRqWqIhk/uXJp0A=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_docker_go_connections",
        importpath = "github.com/docker/go-connections",
        sum = "h1:El9xVISelRB7BuFusrZozjnkIM5YnzCViNKohAFqRJQ=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_docker_go_events",
        importpath = "github.com/docker/go-events",
        sum = "h1:+pKlWGMw7gf6bQ+oDZB4KHQFypsfjYlq/C4rfL7D3g8=",
        version = "v0.0.0-20190806004212-e31b211e4f1c",
    )
    go_repository(
        name = "com_github_docker_go_metrics",
        importpath = "github.com/docker/go-metrics",
        sum = "h1:AgB/0SvBxihN0X8OR4SjsblXkbMvalQ8cjmtKQ2rQV8=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_docker_go_units",
        importpath = "github.com/docker/go-units",
        sum = "h1:69rxXcBk27SvSaaxTtLh/8llcHD8vYHT7WSdRZ/jvr4=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_docker_libtrust",
        importpath = "github.com/docker/libtrust",
        sum = "h1:ZClxb8laGDf5arXfYcAtECDFgAgHklGI8CxgjHnXKJ4=",
        version = "v0.0.0-20150114040149-fa567046d9b1",
    )
    go_repository(
        name = "com_github_docopt_docopt_go",
        importpath = "github.com/docopt/docopt-go",
        sum = "h1:bWDMxwH3px2JBh6AyO7hdCn/PkvCZXii8TGj7sbtEbQ=",
        version = "v0.0.0-20180111231733-ee0de3bc6815",
    )
    go_repository(
        name = "com_github_dustin_go_humanize",
        importpath = "github.com/dustin/go-humanize",
        sum = "h1:GzkhY7T5VNhEkwH0PVJgjz+fX1rhBrR7pRT3mDkpeCY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_dvyukov_go_fuzz",
        importpath = "github.com/dvyukov/go-fuzz",
        sum = "h1:q1oJaUPdmpDm/VyXosjgPgr6wS7c5iV2p0PwJD73bUI=",
        version = "v0.0.0-20210103155950-6a8e9d1f2415",
    )
    go_repository(
        name = "com_github_eapache_go_resiliency",
        importpath = "github.com/eapache/go-resiliency",
        sum = "h1:RRL0nge+cWGlxXbUzJ7yMcq6w2XBEr19dCN6HECGaT0=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_eapache_go_xerial_snappy",
        importpath = "github.com/eapache/go-xerial-snappy",
        sum = "h1:8yY/I9ndfrgrXUbOGObLHKBR4Fl3nZXwM2c7OYTT8hM=",
        version = "v0.0.0-20230111030713-bf00bc1b83b6",
    )
    go_repository(
        name = "com_github_eapache_queue",
        importpath = "github.com/eapache/queue",
        sum = "h1:YOEu7KNc61ntiQlcEeUIoDTJ2o8mQznoNvUhiigpIqc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_ebitengine_purego",
        importpath = "github.com/ebitengine/purego",
        sum = "h1:pNZNC8WofBTN3Nm196An50C5taL/87BhFR/RzKy2o4k=",
        version = "v0.5.0-alpha",
    )
    go_repository(
        name = "com_github_edsrzf_mmap_go",
        importpath = "github.com/edsrzf/mmap-go",
        sum = "h1:CEBF7HpRnUCSJgGUb5h1Gm7e3VkmVDrR8lvWVLtrOFw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_elastic_elastic_transport_go_v8",
        importpath = "github.com/elastic/elastic-transport-go/v8",
        sum = "h1:NeqEz1ty4RQz+TVbUrpSU7pZ48XkzGWQj02k5koahIE=",
        version = "v8.1.0",
    )
    go_repository(
        name = "com_github_elastic_go_elasticsearch_v6",
        importpath = "github.com/elastic/go-elasticsearch/v6",
        sum = "h1:U2HtkBseC1FNBmDr0TR2tKltL6FxoY+niDAlj5M8TK8=",
        version = "v6.8.5",
    )
    go_repository(
        name = "com_github_elastic_go_elasticsearch_v7",
        importpath = "github.com/elastic/go-elasticsearch/v7",
        sum = "h1:49mHcHx7lpCL8cW1aioEwSEVKQF3s+Igi4Ye/QTWwmk=",
        version = "v7.17.1",
    )
    go_repository(
        name = "com_github_elastic_go_elasticsearch_v8",
        importpath = "github.com/elastic/go-elasticsearch/v8",
        sum = "h1:Rn1mcqaIMcNT43hnx2H62cIFZ+B6mjWtzj85BDKrvCE=",
        version = "v8.4.0",
    )
    go_repository(
        name = "com_github_ema_qdisc",
        importpath = "github.com/ema/qdisc",
        sum = "h1:EHLG08FVRbWLg8uRICa3xzC9Zm0m7HyMHfXobWFnXYg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_emicklei_go_restful",
        importpath = "github.com/emicklei/go-restful",
        sum = "h1:rgqiKNjTnFQA6kkhFe16D8epTksy9HQ1MyrbDXSdYhM=",
        version = "v2.16.0+incompatible",
    )
    go_repository(
        name = "com_github_emicklei_go_restful_v3",
        importpath = "github.com/emicklei/go-restful/v3",
        sum = "h1:rAQeMHw1c7zTmncogyy8VvRZwtkmkZ4FxERmMY4rD+g=",
        version = "v3.11.0",
    )
    go_repository(
        name = "com_github_emirpasic_gods",
        importpath = "github.com/emirpasic/gods",
        sum = "h1:QAUIPSaCu4G+POclxeqb3F+WPpdKqFGlw36+yOzGlrg=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_github_envoyproxy_go_control_plane",
        importpath = "github.com/envoyproxy/go-control-plane",
        sum = "h1:4X+VP1GHd1Mhj6IB5mMeGbLCleqxjletLK6K0rbxyZI=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_envoyproxy_protoc_gen_validate",
        importpath = "github.com/envoyproxy/protoc-gen-validate",
        sum = "h1:gVPz/FMfvh57HdSJQyvBtF00j8JU4zdyUgIUNhlgg0A=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_etcd_io_gofail",
        importpath = "github.com/etcd-io/gofail",
        sum = "h1:Y2I0lxOttdUKz+hNaIdG3FtjuQrTmwXun1opRV65IZc=",
        version = "v0.0.0-20190801230047-ad7f989257ca",
    )
    go_repository(
        name = "com_github_euank_go_kmsg_parser",
        importpath = "github.com/euank/go-kmsg-parser",
        sum = "h1:cHD53+PLQuuQyLZeriD1V/esuG4MuU0Pjs5y6iknohY=",
        version = "v2.0.0+incompatible",
    )
    go_repository(
        name = "com_github_evanphx_json_patch",
        importpath = "github.com/evanphx/json-patch",
        sum = "h1:jBYDEEiFBPxA0v50tFdvOzQQTCvpL6mnFh5mB2/l16U=",
        version = "v5.6.0+incompatible",
    )
    go_repository(
        name = "com_github_exponent_io_jsonpath",
        importpath = "github.com/exponent-io/jsonpath",
        sum = "h1:105gxyaGwCFad8crR9dcMQWvV9Hvulu6hwUh4tWPJnM=",
        version = "v0.0.0-20151013193312-d6023ce2651d",
    )
    go_repository(
        name = "com_github_fanliao_go_promise",
        importpath = "github.com/fanliao/go-promise",
        sum = "h1:0eU/faU2oDIB2BkQVM02hgRLJjGzzUuRf19HUhp0394=",
        version = "v0.0.0-20141029170127-1890db352a72",
    )
    go_repository(
        name = "com_github_farsightsec_golang_framestream",
        importpath = "github.com/farsightsec/golang-framestream",
        sum = "h1:/spFQHucTle/ZIPkYqrfshQqPe2VQEzesH243TjIwqA=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_fatih_camelcase",
        importpath = "github.com/fatih/camelcase",
        sum = "h1:hxNvNX/xYBp0ovncs8WyWZrOrpBNub/JfaMvbURyft8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        sum = "h1:8xPHl4/q1VyqGIPif1F+1V3Y3lSmrq01EabUW3CoW5s=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_fatih_structtag",
        importpath = "github.com/fatih/structtag",
        sum = "h1:/OdNE99OxoI/PqaW/SuSK9uxxT3f/tcSZgon/ssNSx4=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_felixge_httpsnoop",
        importpath = "github.com/felixge/httpsnoop",
        sum = "h1:NFTV2Zj1bL4mc9sqWACXbQFVBBg2W3GPvqp8/ESS2Wg=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_flynn_go_docopt",
        importpath = "github.com/flynn/go-docopt",
        sum = "h1:Ss/B3/5wWRh8+emnK0++g5zQzwDTi30W10pKxKc4JXI=",
        version = "v0.0.0-20140912013429-f6dd2ebbb31e",
    )
    go_repository(
        name = "com_github_flynn_go_shlex",
        importpath = "github.com/flynn/go-shlex",
        sum = "h1:BHsljHzVlRcyQhjrss6TZTdY2VfCqZPbv5k3iBFa2ZQ=",
        version = "v0.0.0-20150515145356-3f9db97f8568",
    )
    go_repository(
        name = "com_github_fogleman_gg",
        importpath = "github.com/fogleman/gg",
        sum = "h1:/7zJX8F6AaYQc57WQCyN9cAIz+4bCJGO9B+dyW29am8=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_form3tech_oss_jwt_go",
        importpath = "github.com/form3tech-oss/jwt-go",
        sum = "h1:/l4kBbb4/vGSsdtB5nUe8L7B9mImVMaBPw9L/0TBHU8=",
        version = "v3.2.5+incompatible",
    )
    go_repository(
        name = "com_github_fortytw2_leaktest",
        importpath = "github.com/fortytw2/leaktest",
        sum = "h1:u8491cBMTQ8ft8aeV+adlcytMZylmA5nnwwkRZjI8vw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_francoispqt_gojay",
        importpath = "github.com/francoispqt/gojay",
        sum = "h1:d2m3sFjloqoIUQU3TsHBgj6qg/BVGlTBeHDUmyJnXKk=",
        version = "v1.2.13",
    )
    go_repository(
        name = "com_github_franela_goblin",
        importpath = "github.com/franela/goblin",
        sum = "h1:gb2Z18BhTPJPpLQWj4T+rfKHYCHxRHCtRxhKKjRidVw=",
        version = "v0.0.0-20200105215937-c9ffbefa60db",
    )
    go_repository(
        name = "com_github_franela_goreq",
        importpath = "github.com/franela/goreq",
        sum = "h1:a9ENSRDFBUPkJ5lCgVZh26+ZbGyoVJG7yb5SSzF5H54=",
        version = "v0.0.0-20171204163338-bcd34c9993f8",
    )
    go_repository(
        name = "com_github_frankban_quicktest",
        importpath = "github.com/frankban/quicktest",
        sum = "h1:dfYrrRyLtiqT9GyKXgdh+k4inNeTvmGbuSgZ3lx3GhA=",
        version = "v1.14.5",
    )
    go_repository(
        name = "com_github_fsnotify_fsnotify",
        importpath = "github.com/fsnotify/fsnotify",
        sum = "h1:8JEhPFa5W2WU7YfeZzPNqzMP6Lwt7L2715Ggo0nosvA=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_fsouza_fake_gcs_server",
        importpath = "github.com/fsouza/fake-gcs-server",
        sum = "h1:OeH75kBZcZa3ZE+zz/mFdJ2btt9FgqfjI7gIh9+5fvk=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_github_fullsailor_pkcs7",
        importpath = "github.com/fullsailor/pkcs7",
        sum = "h1:RDBNVkRviHZtvDvId8XSGPu3rmpmSe+wKRcEWNgsfWU=",
        version = "v0.0.0-20190404230743-d7302db945fa",
    )
    go_repository(
        name = "com_github_fullstorydev_grpcurl",
        importpath = "github.com/fullstorydev/grpcurl",
        sum = "h1:2II5e++aFnctnPJir3GL6cPSwF69Ord1u/9O+fv1vrI=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_fvbommel_sortorder",
        importpath = "github.com/fvbommel/sortorder",
        sum = "h1:fUmoe+HLsBTctBDoaBwpQo5N+nrCp8g/BjKb/6ZQmYw=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_fxamacker_cbor_v2",
        importpath = "github.com/fxamacker/cbor/v2",
        sum = "h1:sU6J2usfADwWlYDAFhZBQ6TnLFBHxgesMrQfQgk1tWA=",
        version = "v2.6.0",
    )
    go_repository(
        name = "com_github_gabriel_vasile_mimetype",
        importpath = "github.com/gabriel-vasile/mimetype",
        sum = "h1:w5qFW6JKBz9Y393Y4q372O9A7cUSequkh1Q7OhCmWKU=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_garyburd_redigo",
        importpath = "github.com/garyburd/redigo",
        sum = "h1:HCeeRluvAgMusMomi1+6Y5dmFOdYV/JzoRrrbFlkGIc=",
        version = "v1.6.3",
    )
    go_repository(
        name = "com_github_getsentry_raven_go",
        importpath = "github.com/getsentry/raven-go",
        sum = "h1:no+xWJRb5ZI7eE8TWgIq1jLulQiIoLG0IfYxv5JYMGs=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gin_contrib_sse",
        importpath = "github.com/gin-contrib/sse",
        sum = "h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gin_gonic_gin",
        importpath = "github.com/gin-gonic/gin",
        sum = "h1:4idEAncQnU5cB7BeOkPtxjfCSye0AAm1R0RVIqJ+Jmg=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_gliderlabs_ssh",
        importpath = "github.com/gliderlabs/ssh",
        sum = "h1:6zsha5zo/TWhRhwqCD3+EarCAgZ2yN28ipRnGPnwkI0=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_globalsign_mgo",
        importpath = "github.com/globalsign/mgo",
        sum = "h1:DujepqpGd1hyOd7aW59XpK7Qymp8iy83xq74fLr21is=",
        version = "v0.0.0-20181015135952-eeefdecb41b8",
    )
    go_repository(
        name = "com_github_go_chi_chi",
        importpath = "github.com/go-chi/chi",
        sum = "h1:2ZcJZozJ+rj6BA0c19ffBUGXEKAT/aOLOtQjD46vBRA=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_go_chi_chi_v5",
        importpath = "github.com/go-chi/chi/v5",
        sum = "h1:DBPx88FjZJH3FsICfDAfIfnb7XxKIYVGG6lOPlhENAg=",
        version = "v5.0.0",
    )
    go_repository(
        name = "com_github_go_delve_delve",
        importpath = "github.com/go-delve/delve",
        patch_args = ["-p1"],
        # https://github.com/bazelbuild/bazel-gazelle/issues/1424#issuecomment-1568744161
        patch_tool = "patch",
        patches = [
            "//third_party/go/patches:delve-fix-cgo.patch",
        ],
        sum = "h1:gsRTPR3Yi61RpeuCFvJb6vIxB3xABx6pnNKGISxdsSU=",
        version = "v1.8.2",
    )
    go_repository(
        name = "com_github_go_delve_liner",
        importpath = "github.com/go-delve/liner",
        sum = "h1:0hGpZh6vjI6LFTlXuHFCX9PebaluzDzps2owdMrrSuk=",
        version = "v1.2.2-1",
    )
    go_repository(
        name = "com_github_go_errors_errors",
        importpath = "github.com/go-errors/errors",
        sum = "h1:J6MZopCL4uSllY1OfXM374weqZFFItUbrImctkmUxIA=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_go_fonts_dejavu",
        importpath = "github.com/go-fonts/dejavu",
        sum = "h1:JSajPXURYqpr+Cu8U9bt8K+XcACIHWqWrvWCKyeFmVQ=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_go_fonts_latin_modern",
        importpath = "github.com/go-fonts/latin-modern",
        sum = "h1:5/Tv1Ek/QCr20C6ZOz15vw3g7GELYL98KWr8Hgo+3vk=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_fonts_liberation",
        importpath = "github.com/go-fonts/liberation",
        sum = "h1:jAkAWJP4S+OsrPLZM4/eC9iW7CtHy+HBXrEwZXWo5VM=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_fonts_stix",
        importpath = "github.com/go-fonts/stix",
        sum = "h1:UlZlgrvvmT/58o573ot7NFw0vZasZ5I6bcIft/oMdgg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_go_ini_ini",
        importpath = "github.com/go-ini/ini",
        sum = "h1:Mujh4R/dH6YL8bxuISne3xX2+qcQ9p0IxKAP6ExWoUo=",
        version = "v1.25.4",
    )
    go_repository(
        name = "com_github_go_kit_kit",
        importpath = "github.com/go-kit/kit",
        sum = "h1:dXFJfIHVvUcpSgDOV+Ne6t7jXri8Tfv2uOLHUZ2XNuo=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_go_kit_log",
        importpath = "github.com/go-kit/log",
        sum = "h1:MRVx0/zhvdseW+Gza6N9rVzU/IVzaeE1SFI4raAhmBU=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_go_latex_latex",
        importpath = "github.com/go-latex/latex",
        sum = "h1:6zl3BbBhdnMkpSj2YY30qV3gDcVBGtFgVsV3+/i+mKQ=",
        version = "v0.0.0-20210823091927-c0d11ff05a81",
    )
    go_repository(
        name = "com_github_go_logfmt_logfmt",
        importpath = "github.com/go-logfmt/logfmt",
        sum = "h1:otpy5pqBCBZ1ng9RQ0dPu4PN7ba75Y/aA+UpowDyNVA=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_go_logr_logr",
        importpath = "github.com/go-logr/logr",
        sum = "h1:pKouT5E8xu9zeFC39JXRDukb6JFQPXM5p5I91188VAQ=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_go_logr_stdr",
        importpath = "github.com/go-logr/stdr",
        sum = "h1:hSWxHoqTgW2S2qGc0LTAI563KZ5YKYRhT3MFKZMbjag=",
        version = "v1.2.2",
    )
    go_repository(
        name = "com_github_go_logr_zapr",
        importpath = "github.com/go-logr/zapr",
        sum = "h1:XGdV8XW8zdwFiwOA2Dryh1gj2KRQyOOoNmBy4EplIcQ=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_go_openapi_jsonpointer",
        importpath = "github.com/go-openapi/jsonpointer",
        sum = "h1:eCs3fxoIi3Wh6vtgmLTOjdhSpiqphQ+DaPn38N2ZdrE=",
        version = "v0.19.6",
    )
    go_repository(
        name = "com_github_go_openapi_jsonreference",
        importpath = "github.com/go-openapi/jsonreference",
        sum = "h1:3sVjiK66+uXK/6oQ8xgcRKcFgQ5KXa2KvnJRumpMGbE=",
        version = "v0.20.2",
    )
    go_repository(
        name = "com_github_go_openapi_swag",
        importpath = "github.com/go-openapi/swag",
        sum = "h1:yMBqmnQ0gyZvEb/+KzuWZOXgllrXT4SADYbvDaXHv/g=",
        version = "v0.22.3",
    )
    go_repository(
        name = "com_github_go_pdf_fpdf",
        importpath = "github.com/go-pdf/fpdf",
        sum = "h1:MlgtGIfsdMEEQJr2le6b/HNr1ZlQwxyWr77r2aj2U/8=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_go_pg_pg_v10",
        importpath = "github.com/go-pg/pg/v10",
        sum = "h1:CMKJqLgTrfpE/aOVeLdybezR2om071Vh38OLZjsyMI0=",
        version = "v10.11.0",
    )
    go_repository(
        name = "com_github_go_pg_zerochecker",
        importpath = "github.com/go-pg/zerochecker",
        sum = "h1:pp7f72c3DobMWOb2ErtZsnrPaSvHd2W4o9//8HtF4mU=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_go_playground_assert_v2",
        importpath = "github.com/go-playground/assert/v2",
        sum = "h1:MsBgLAaY856+nPRTKrp3/OZK38U/wa0CcBYNjji3q3A=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_go_playground_locales",
        importpath = "github.com/go-playground/locales",
        sum = "h1:EWaQ/wswjilfKLTECiXz7Rh+3BjFhfDFKv/oXslEjJA=",
        version = "v0.14.1",
    )
    go_repository(
        name = "com_github_go_playground_universal_translator",
        importpath = "github.com/go-playground/universal-translator",
        sum = "h1:Bcnm0ZwsGyWbCzImXv+pAJnYK9S473LQFuzCbDbfSFY=",
        version = "v0.18.1",
    )
    go_repository(
        name = "com_github_go_playground_validator_v10",
        importpath = "github.com/go-playground/validator/v10",
        sum = "h1:vgvQWe3XCz3gIeFDm/HnTIbj6UGmg/+t63MyGU2n5js=",
        version = "v10.14.0",
    )
    go_repository(
        name = "com_github_go_redis_redis",
        importpath = "github.com/go-redis/redis",
        sum = "h1:K0pv1D7EQUjfyoMql+r/jZqCLizCGKFlFgcHWWmHQjg=",
        version = "v6.15.9+incompatible",
    )
    go_repository(
        name = "com_github_go_redis_redis_v7",
        importpath = "github.com/go-redis/redis/v7",
        sum = "h1:I4C4a8UGbFejiVjtYVTRVOiMIJ5pm5Yru6ibvDX/OS0=",
        version = "v7.1.0",
    )
    go_repository(
        name = "com_github_go_redis_redis_v8",
        importpath = "github.com/go-redis/redis/v8",
        sum = "h1:AcZZR7igkdvfVmQTPnu9WE37LRrO/YrBH5zWyjDC0oI=",
        version = "v8.11.5",
    )
    go_repository(
        name = "com_github_go_sql_driver_mysql",
        importpath = "github.com/go-sql-driver/mysql",
        sum = "h1:lUIinVbN1DY0xBg0eMOzmmtGoHwWBbvnWubQUrtU8EI=",
        version = "v1.7.1",
    )
    go_repository(
        name = "com_github_go_stack_stack",
        importpath = "github.com/go-stack/stack",
        sum = "h1:5SgMzNM5HxrEjV0ww2lTmX6E2Izsfxas4+YHWRs3Lsk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_go_task_slim_sprig",
        importpath = "github.com/go-task/slim-sprig",
        sum = "h1:tfuBGBXKqDEevZMzYi5KSi8KkcZtzBcTgAUUtapy0OI=",
        version = "v0.0.0-20230315185526-52ccab3ef572",
    )
    go_repository(
        name = "com_github_gobuffalo_attrs",
        importpath = "github.com/gobuffalo/attrs",
        sum = "h1:hSkbZ9XSyjyBirMeqSqUrK+9HboWrweVlzRNqoBi2d4=",
        version = "v0.0.0-20190224210810-a9411de4debd",
    )
    go_repository(
        name = "com_github_gobuffalo_depgen",
        importpath = "github.com/gobuffalo/depgen",
        sum = "h1:31atYa/UW9V5q8vMJ+W6wd64OaaTHUrCUXER358zLM4=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobuffalo_envy",
        importpath = "github.com/gobuffalo/envy",
        sum = "h1:GlXgaiBkmrYMHco6t4j7SacKO4XUjvh5pwXh0f4uxXU=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_gobuffalo_flect",
        importpath = "github.com/gobuffalo/flect",
        sum = "h1:3GQ53z7E3o00C/yy7Ko8VXqQXoJGLkrTQCLTF1EjoXU=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_gobuffalo_genny",
        importpath = "github.com/gobuffalo/genny",
        sum = "h1:iQ0D6SpNXIxu52WESsD+KoQ7af2e3nCfnSBoSF/hKe0=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_gobuffalo_gitgen",
        importpath = "github.com/gobuffalo/gitgen",
        sum = "h1:mSVZ4vj4khv+oThUfS+SQU3UuFIZ5Zo6UNcvK8E8Mz8=",
        version = "v0.0.0-20190315122116-cc086187d211",
    )
    go_repository(
        name = "com_github_gobuffalo_gogen",
        importpath = "github.com/gobuffalo/gogen",
        sum = "h1:dLg+zb+uOyd/mKeQUYIbwbNmfRsr9hd/WtYWepmayhI=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_gobuffalo_here",
        importpath = "github.com/gobuffalo/here",
        sum = "h1:hYrd0a6gDmWxBM4TnrGw8mQg24iSVoIkHEk7FodQcBI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_gobuffalo_logger",
        importpath = "github.com/gobuffalo/logger",
        sum = "h1:8thhT+kUJMTMy3HlX4+y9Da+BNJck+p109tqqKp7WDs=",
        version = "v0.0.0-20190315122211-86e12af44bc2",
    )
    go_repository(
        name = "com_github_gobuffalo_mapi",
        importpath = "github.com/gobuffalo/mapi",
        sum = "h1:fq9WcL1BYrm36SzK6+aAnZ8hcp+SrmnDyAxhNx8dvJk=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_gobuffalo_packd",
        importpath = "github.com/gobuffalo/packd",
        sum = "h1:4sGKOD8yaYJ+dek1FDkwcxCHA40M4kfKgFHx8N2kwbU=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_gobuffalo_packr_v2",
        importpath = "github.com/gobuffalo/packr/v2",
        sum = "h1:Ir9W9XIm9j7bhhkKE9cokvtTl1vBm62A/fene/ZCj6A=",
        version = "v2.2.0",
    )
    go_repository(
        name = "com_github_gobuffalo_syncx",
        importpath = "github.com/gobuffalo/syncx",
        sum = "h1:tpom+2CJmpzAWj5/VEHync2rJGi+epHNIeRSWjzGA+4=",
        version = "v0.0.0-20190224160051-33c29581e754",
    )
    go_repository(
        name = "com_github_gobwas_glob",
        importpath = "github.com/gobwas/glob",
        sum = "h1:A4xDbljILXROh+kObIiy5kIaPYD8e96x1tgBhUI5J+Y=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_github_gobwas_httphead",
        importpath = "github.com/gobwas/httphead",
        sum = "h1:s+21KNqlpePfkah2I+gwHF8xmJWRjooY+5248k6m4A0=",
        version = "v0.0.0-20180130184737-2c6c146eadee",
    )
    go_repository(
        name = "com_github_gobwas_pool",
        importpath = "github.com/gobwas/pool",
        sum = "h1:QEmUOlnSjWtnpRGHF3SauEiOsy82Cup83Vf2LcMlnc8=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_gobwas_ws",
        importpath = "github.com/gobwas/ws",
        sum = "h1:CoAavW/wd/kulfZmSIBt6p24n4j7tHgNVCjsfHVNUbo=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_goccy_go_json",
        importpath = "github.com/goccy/go-json",
        sum = "h1:CrxCmQqYDkv1z7lO7Wbh2HN93uovUHgrECaO5ZrCXAU=",
        version = "v0.10.2",
    )
    go_repository(
        name = "com_github_gocql_gocql",
        importpath = "github.com/gocql/gocql",
        sum = "h1:6ImvI6U901e1ezn/8u2z3bh1DZIvMOia0yTSBxhy4Ao=",
        version = "v0.0.0-20220224095938-0eacd3183625",
    )
    go_repository(
        name = "com_github_godbus_dbus",
        importpath = "github.com/godbus/dbus",
        sum = "h1:BWhy2j3IXJhjCbC68FptL43tDKIq8FladmaTs3Xs7Z8=",
        version = "v0.0.0-20190422162347-ade71ed3457e",
    )
    go_repository(
        name = "com_github_godbus_dbus_v5",
        importpath = "github.com/godbus/dbus/v5",
        sum = "h1:4KLkAxT3aOY8Li4FRJe/KvhoNFFxo0m6fNuFUO8QJUk=",
        version = "v5.1.0",
    )
    go_repository(
        name = "com_github_gofiber_fiber_v2",
        importpath = "github.com/gofiber/fiber/v2",
        sum = "h1:yit3E4kHf178B60p5CQBa/3v+WVuziWMa/G2ZNyLJB0=",
        version = "v2.43.0",
    )
    go_repository(
        name = "com_github_gofrs_flock",
        importpath = "github.com/gofrs/flock",
        sum = "h1:+gYjHKf32LDeiEEFhQaotPbLuUXjY5ZqxKgXy7n59aw=",
        version = "v0.8.1",
    )
    go_repository(
        name = "com_github_gofrs_uuid",
        importpath = "github.com/gofrs/uuid",
        sum = "h1:1SD/1F5pU8p29ybwgQSwpQk+mwdRrXCYuPhW6m+TnJw=",
        version = "v4.0.0+incompatible",
    )
    go_repository(
        name = "com_github_gogo_googleapis",
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/googleapis",
        sum = "h1:1Yx4Myt7BxzvUr5ldGSbwYiZG6t9wGBZ+8/fX3Wvtq0=",
        version = "v1.4.1",
    )
    go_repository(
        name = "com_github_gogo_protobuf",
        build_file_proto_mode = "disable",
        importpath = "github.com/gogo/protobuf",
        sum = "h1:Ov1cvc58UF3b5XjBnZv7+opcTcQFZebYjWzi34vdm4Q=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_golang_freetype",
        importpath = "github.com/golang/freetype",
        sum = "h1:DACJavvAHhabrF08vX0COfcOBJRhZ8lUbR+ZWIs0Y5g=",
        version = "v0.0.0-20170609003504-e2365dfdc4a0",
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:uCdmnmatrKCgMBlM4rMuJZWOkPDqdbZPnrMXDY4gI68=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_golang_groupcache",
        importpath = "github.com/golang/groupcache",
        sum = "h1:oI5xCqsCo564l8iNU+DwB5epxmsaqB+rhGL0m5jtYqE=",
        version = "v0.0.0-20210331224755-41bb18bfe9da",
    )
    go_repository(
        name = "com_github_golang_jwt_jwt_v4",
        importpath = "github.com/golang-jwt/jwt/v4",
        sum = "h1:7cYmW1XlMY7h7ii7UhUyChSgS5wUJEnm9uZVTGqOWzg=",
        version = "v4.5.0",
    )
    go_repository(
        name = "com_github_golang_migrate_migrate_v4",
        importpath = "github.com/golang-migrate/migrate/v4",
        sum = "h1:vU+M05vs6jWHKDdmE1Ecwj0BznygFc4QsdRe2E/L7kc=",
        version = "v4.15.2",
    )
    go_repository(
        name = "com_github_golang_mock",
        importpath = "github.com/golang/mock",
        sum = "h1:ErTB+efbowRARo13NNdxyJji2egdxLGQhRaY+DUumQc=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_golang_protobuf",
        importpath = "github.com/golang/protobuf",
        sum = "h1:i7eJL8qZTpSEXOPTxNKhASYpMn+8e5Q6AdndVa1dWek=",
        version = "v1.5.4",
    )
    go_repository(
        name = "com_github_golang_snappy",
        importpath = "github.com/golang/snappy",
        sum = "h1:yAGX7huGHXlcLOEtBnF4w7FQwA26wojNCwOYAEhLjQM=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_golang_sql_civil",
        importpath = "github.com/golang-sql/civil",
        sum = "h1:au07oEsX2xN0ktxqI+Sida1w446QrXBRJ0nee3SNZlA=",
        version = "v0.0.0-20220223132316-b832511892a9",
    )
    go_repository(
        name = "com_github_golang_sql_sqlexp",
        importpath = "github.com/golang-sql/sqlexp",
        sum = "h1:ZCD6MBpcuOVfGVqsEmY5/4FtYiKz6tSyUv9LPEDei6A=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_golangplus_bytes",
        importpath = "github.com/golangplus/bytes",
        sum = "h1:YQKBijBVMsBxIiXT4IEhlKR2zHohjEqPole4umyDX+c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_golangplus_fmt",
        importpath = "github.com/golangplus/fmt",
        sum = "h1:FnUKtw86lXIPfBMc3FimNF3+ABcV+aH5F17OOitTN+E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_golangplus_testing",
        importpath = "github.com/golangplus/testing",
        sum = "h1:+ZeeiKZENNOMkTTELoSySazi+XaEhVO0mb+eanrSEUQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_gomodule_redigo",
        importpath = "github.com/gomodule/redigo",
        sum = "h1:Sl3u+2BI/kk+VEatbj0scLdrFhjPmbxOc1myhDP41ws=",
        version = "v1.8.9",
    )
    go_repository(
        name = "com_github_google_btree",
        importpath = "github.com/google/btree",
        sum = "h1:gK4Kx5IaGY9CD5sPJ36FHiBJ6ZXl0kilRiiCj+jdYp4=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_google_cadvisor",
        build_file_proto_mode = "disable_global",  # keep
        importpath = "github.com/google/cadvisor",
        sum = "h1:1PYeiORXmcFYi609M4Qvq5IzcvcVaWgYxDt78uH8jYA=",
        version = "v0.49.0",
    )
    go_repository(
        name = "com_github_google_cel_go",
        # Gazelle needs to regenerate the upstream build files due to a
        # mismatch in how the Google APIs repository is canonically called.
        # See https://github.com/bazelbuild/rules_go/issues/1986 for more
        # details.
        build_file_generation = "on",
        build_naming_convention = "go_default_library",
        importpath = "github.com/google/cel-go",
        replace = "github.com/google/cel-go",
        sum = "h1:j9m730pMZt1Fc4oKhCLUHfjj6527LuhYcYw0Rl8gqto=",
        version = "v0.17.8",
    )
    go_repository(
        name = "com_github_google_certificate_transparency_go",
        importpath = "github.com/google/certificate-transparency-go",
        sum = "h1:4hE0GEId6NAW28dFpC+LrRGwQX5dtmXQGDbg8+/MZOM=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_google_flatbuffers",
        importpath = "github.com/google/flatbuffers",
        sum = "h1:ivUb1cGomAB101ZM1T0nOiWz9pSrTMoa9+EiY7igmkM=",
        version = "v2.0.8+incompatible",
    )
    go_repository(
        name = "com_github_google_gnostic",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gnostic",
        sum = "h1:FhTMOKj2VhjpouxvWJAV1TL304uMlb9zcDqkl6cEI54=",
        version = "v0.5.7-v3refs",
    )
    go_repository(
        name = "com_github_google_gnostic_models",
        build_file_proto_mode = "disable",
        importpath = "github.com/google/gnostic-models",
        sum = "h1:yo/ABAfM5IMRsS1VnXjTBvUb61tFIHozhlYvRgGre9I=",
        version = "v0.6.8",
    )
    go_repository(
        name = "com_github_google_go_attestation",
        importpath = "github.com/google/go-attestation",
        sum = "h1:hHhPfym1TZm88L7sWmdc/moikHt80ls6mEiU+QvhRvk=",
        version = "v0.4.3",
    )
    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_go_containerregistry",
        importpath = "github.com/google/go-containerregistry",
        sum = "h1:z58vMqHxuwvAsVwvKEkmVBz2TlgBgH5k6koEXBtlYkw=",
        version = "v0.14.0",
    )
    go_repository(
        name = "com_github_google_go_dap",
        importpath = "github.com/google/go-dap",
        sum = "h1:Y1RHGUtv3R8y6sXq2dtGRMYrFB2hSqyFVws7jucrzX4=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_github_google_go_github_v28",
        importpath = "github.com/google/go-github/v28",
        sum = "h1:kORf5ekX5qwXO2mGzXXOjMe/g6ap8ahVe0sBEulhSxo=",
        version = "v28.1.1",
    )
    go_repository(
        name = "com_github_google_go_github_v39",
        importpath = "github.com/google/go-github/v39",
        sum = "h1:rNNM311XtPOz5rDdsJXAp2o8F67X9FnROXTvto3aSnQ=",
        version = "v39.2.0",
    )
    go_repository(
        name = "com_github_google_go_licenses",
        importpath = "github.com/google/go-licenses",
        sum = "h1:JtmsUf+m+KdwCOgLG578T0Mvd0+l+dezPrJh5KYnXZg=",
        version = "v0.0.0-20210329231322-ce1d9163b77d",
    )
    go_repository(
        name = "com_github_google_go_pkcs11",
        importpath = "github.com/google/go-pkcs11",
        sum = "h1:OF1IPgv+F4NmqmJ98KTjdN97Vs1JxDPB3vbmYzV2dpk=",
        version = "v0.2.1-0.20230907215043-c6f79328ddf9",
    )
    go_repository(
        name = "com_github_google_go_querystring",
        importpath = "github.com/google/go-querystring",
        sum = "h1:AnCroh3fv4ZBgVIf1Iwtovgjaw/GiKJo8M8yD/fhyJ8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_google_go_replayers_grpcreplay",
        importpath = "github.com/google/go-replayers/grpcreplay",
        sum = "h1:eNb1y9rZFmY4ax45uEEECSa8fsxGRU+8Bil52ASAwic=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_go_replayers_httpreplay",
        importpath = "github.com/google/go-replayers/httpreplay",
        sum = "h1:AX7FUb4BjrrzNvblr/OlgwrmFiep6soj5K2QSDW7BGk=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_go_tpm",
        importpath = "github.com/google/go-tpm",
        sum = "h1:P/ZFNBZYXRxc+z7i5uyd8VP7MaDteuLZInzrH2idRGo=",
        version = "v0.3.3",
    )
    go_repository(
        name = "com_github_google_go_tpm_tools",
        importpath = "github.com/google/go-tpm-tools",
        patch_args = ["-p1"],
        pre_patches = [
            "//third_party/go/patches:go-tpm-tools-fix-proto-paths.patch",
        ],
        sum = "h1:gcjAoa/iK98ybY41Icn2zkGlJLhVg44nOmI6fFju5IY=",
        version = "v0.3.5",
    )
    go_repository(
        name = "com_github_google_go_tspi",
        importpath = "github.com/google/go-tspi",
        sum = "h1:LnpS22S8V1HqbxjveESGAazHhi6BX9SwI2Rij7qZcXQ=",
        version = "v0.2.1-0.20190423175329-115dea689aad",
    )
    go_repository(
        name = "com_github_google_gofuzz",
        importpath = "github.com/google/gofuzz",
        sum = "h1:xRy4A+RhZaiKjJ1bPfwQ8sedCA+YS2YcCHW6ec7JMi0=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_google_gopacket",
        importpath = "github.com/google/gopacket",
        sum = "h1:ves8RnFZPGiFnTS0uPQStjwru6uO6h+nlr9j6fL7kF8=",
        version = "v1.1.19",
    )
    go_repository(
        name = "com_github_google_licenseclassifier",
        importpath = "github.com/google/licenseclassifier",
        sum = "h1:EfzlPF5MRmoWsCGvSkPZ1Nh9uVzHf4FfGnDQ6CXd2NA=",
        version = "v0.0.0-20210325184830-bb04aff29e72",
    )
    go_repository(
        name = "com_github_google_martian",
        importpath = "github.com/google/martian",
        sum = "h1:xmapqc1AyLoB+ddYT6r04bD9lIjlOqGaREovi0SzFaE=",
        version = "v2.1.1-0.20190517191504-25dcb96d9e51+incompatible",
    )
    go_repository(
        name = "com_github_google_martian_v3",
        importpath = "github.com/google/martian/v3",
        sum = "h1:IqNFLAmvJOgVlpdEBiQbDc2EwKW77amAycfTuWKdfvw=",
        version = "v3.3.2",
    )
    go_repository(
        name = "com_github_google_nftables",
        importpath = "github.com/google/nftables",
        sum = "h1:UpwO4s4fHzquJ3SjQpZhcg1fga6RA4PsrKBGOj/ilvg=",
        version = "v0.0.0-20220221214239-211824995dcb",
    )
    go_repository(
        name = "com_github_google_pprof",
        importpath = "github.com/google/pprof",
        sum = "h1:PEOGDI1kkyW37YqPWHLHc+D20D9+87Wt12TCcfTUo5Q=",
        version = "v0.0.0-20230509042627-b1315fad0c5a",
    )
    go_repository(
        name = "com_github_google_renameio",
        importpath = "github.com/google/renameio",
        sum = "h1:GOZbcHa3HfsPKPlmyPyN2KEohoMXOhdMbHrvbpl2QaA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_google_rpmpack",
        importpath = "github.com/google/rpmpack",
        sum = "h1:BW6OvS3kpT5UEPbCZ+KyX/OB4Ks9/MNMhWjqPPkZxsE=",
        version = "v0.0.0-20191226140753-aa36bfddb3a0",
    )
    go_repository(
        name = "com_github_google_s2a_go",
        importpath = "github.com/google/s2a-go",
        sum = "h1:60BLSyTrOV4/haCDW4zb1guZItoSq8foHCXrAnjBo/o=",
        version = "v0.1.7",
    )
    go_repository(
        name = "com_github_google_shlex",
        importpath = "github.com/google/shlex",
        sum = "h1:El6M4kTTCOh6aBiKaUGG7oYTSPP8MxqL4YI3kZKwcP4=",
        version = "v0.0.0-20191202100458-e7afc7fbc510",
    )
    go_repository(
        name = "com_github_google_subcommands",
        importpath = "github.com/google/subcommands",
        sum = "h1:8nlgEAjIalk6uj/CGKCdOO8CQqTeysvcW4RFZ6HbkGM=",
        version = "v1.0.2-0.20190508160503-636abe8753b8",
    )
    go_repository(
        name = "com_github_google_trillian",
        importpath = "github.com/google/trillian",
        sum = "h1:Wa7XHCVzl8RLsUOr2SzoHUZHYjv0G8KMO1xZGamYkbA=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        sum = "h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_google_wire",
        importpath = "github.com/google/wire",
        sum = "h1:imGQZGEVEHpje5056+K+cgdO72p0LQv2xIIFXNGUf60=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_googleapis_enterprise_certificate_proxy",
        importpath = "github.com/googleapis/enterprise-certificate-proxy",
        sum = "h1:Vie5ybvEvT75RniqhfFxPRy3Bf7vr3h0cechB90XaQs=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_googleapis_gax_go",
        importpath = "github.com/googleapis/gax-go",
        sum = "h1:silFMLAnr330+NRuag/VjIGF7TLp/LBrV2CJKFLWEww=",
        version = "v2.0.2+incompatible",
    )
    go_repository(
        name = "com_github_googleapis_gax_go_v2",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gax-go/v2",
        sum = "h1:A+gCJKdRfqXkr+BIRGtZLibNXf0m1f9E4HG56etFpas=",
        version = "v2.12.0",
    )
    go_repository(
        name = "com_github_googleapis_gnostic",
        build_file_proto_mode = "disable",
        importpath = "github.com/googleapis/gnostic",
        sum = "h1:9fHAtK0uDfpveeqqo1hkEZJcFvYXAiCN3UutL8F9xHw=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_github_googleapis_go_type_adapters",
        importpath = "github.com/googleapis/go-type-adapters",
        sum = "h1:9XdMn+d/G57qq1s8dNc5IesGCXHf6V2HZ2JwRxfA2tA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_googleapis_google_cloud_go_testing",
        importpath = "github.com/googleapis/google-cloud-go-testing",
        sum = "h1:tlyzajkF3030q6M8SvmJSemC9DTHL/xaMa18b65+JM4=",
        version = "v0.0.0-20200911160855-bcd43fbb19e8",
    )
    go_repository(
        name = "com_github_googlecloudplatform_cloudsql_proxy",
        importpath = "github.com/GoogleCloudPlatform/cloudsql-proxy",
        sum = "h1:sTOp2Ajiew5XIH92YSdwhYc+bgpUX5j5TKK/Ac8Saw8=",
        version = "v0.0.0-20191009163259-e802c2cb94ae",
    )
    go_repository(
        name = "com_github_googlecloudplatform_k8s_cloud_provider",
        importpath = "github.com/GoogleCloudPlatform/k8s-cloud-provider",
        sum = "h1:Heo1J/ttaQFgGJSVnCZquy3e5eH5j1nqxBuomztB3P0=",
        version = "v1.18.1-0.20220218231025-f11817397a1b",
    )
    go_repository(
        name = "com_github_gopherjs_gopherjs",
        importpath = "github.com/gopherjs/gopherjs",
        sum = "h1:EGx4pi6eqNxGaHF6qqu48+N2wcFQ5qg5FXgOdqsJ5d8=",
        version = "v0.0.0-20181017120253-0766667cb4d1",
    )
    go_repository(
        name = "com_github_gordonklaus_ineffassign",
        importpath = "github.com/gordonklaus/ineffassign",
        sum = "h1:vc7Dmrk4JwS0ZPS6WZvWlwDflgDTA26jItmbSj83nug=",
        version = "v0.0.0-20200309095847-7953dde2c7bf",
    )
    go_repository(
        name = "com_github_goreleaser_goreleaser",
        importpath = "github.com/goreleaser/goreleaser",
        sum = "h1:3ua5fSYvc/doKd5sVTzvXbyqW6H/eyuvy46woLrUFic=",
        version = "v0.134.0",
    )
    go_repository(
        name = "com_github_goreleaser_nfpm",
        importpath = "github.com/goreleaser/nfpm",
        sum = "h1:AEnu9XVmupRDTR930Z2rAs31Mj6sLIPxFcR9ESYvgDA=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_gorilla_context",
        importpath = "github.com/gorilla/context",
        sum = "h1:AWwleXJkX/nhcU9bZSnZoi3h/qGYqQAGhq6zZe/aQW8=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_gorilla_handlers",
        importpath = "github.com/gorilla/handlers",
        sum = "h1:0QniY0USkHQ1RGCLfKxeNHK9bkDHGRYGNDFBCS+YARg=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_gorilla_mux",
        importpath = "github.com/gorilla/mux",
        sum = "h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvkdNIeFDP5koI=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_gorilla_websocket",
        importpath = "github.com/gorilla/websocket",
        sum = "h1:PPwGk2jz7EePpoHN/+ClbZu8SPxiqlu12wZP/3sWmnc=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_graph_gophers_graphql_go",
        importpath = "github.com/graph-gophers/graphql-go",
        sum = "h1:Eb9x/q6MFpCLz7jBCiP/WTxjSDrYLR1QY41SORZyNJ0=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_gregjones_httpcache",
        importpath = "github.com/gregjones/httpcache",
        sum = "h1:pdN6V1QBWetyv/0+wjACpqVH+eVULgEjkurDLq3goeM=",
        version = "v0.0.0-20180305231024-9cad4c3443a7",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        sum = "h1:+9834+KizmvFV7pXQGSXQTsaWhq2GjuNUt0aUU0YBYw=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_prometheus",
        importpath = "github.com/grpc-ecosystem/go-grpc-prometheus",
        sum = "h1:Ovs26xHkKqVztRpIrF/92BcuyuQ/YW4NSIpoGtfXNho=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        # Force generation of new build files for this package, required due to changes in how gazelle handles @go_googleapis.
        build_file_generation = "on",
        build_naming_convention = "go_default_library",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        sum = "h1:gmcG1KaJ57LophUzW0Hy8NmPhnMZb4M0+kPpLofRdBo=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway_v2",
        # Force generation of new build files for this package, required due to changes in how gazelle handles @go_googleapis.
        build_file_generation = "on",
        build_naming_convention = "go_default_library",
        importpath = "github.com/grpc-ecosystem/grpc-gateway/v2",
        sum = "h1:YBftPWNWd4WwGqtY2yeZL2ef8rHAxPBD8KFhJpmcqms=",
        version = "v2.16.0",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_grpc_opentracing",
        importpath = "github.com/grpc-ecosystem/grpc-opentracing",
        sum = "h1:MJG/KsmcqMwFAkh8mTnAwhyKoB+sTAnY4CACC110tbU=",
        version = "v0.0.0-20180507213350-8e809c8a8645",
    )
    go_repository(
        name = "com_github_hailocab_go_hostpool",
        importpath = "github.com/hailocab/go-hostpool",
        sum = "h1:5upAirOpQc1Q53c0bnx2ufif5kANL7bfZWcc6VJWJd8=",
        version = "v0.0.0-20160125115350-e80d13ce29ed",
    )
    go_repository(
        name = "com_github_hanwen_go_fuse_v2",
        importpath = "github.com/hanwen/go-fuse/v2",
        sum = "h1:t5ivNIH2PK+zw4OBul/iJjsoG9K6kXo4nMDoBpciC8A=",
        version = "v2.3.0",
    )
    go_repository(
        name = "com_github_hashicorp_consul_api",
        importpath = "github.com/hashicorp/consul/api",
        sum = "h1:HXNYlRkkM/t+Y/Yhxtwcy02dlYwIaoxzvxPnS+cqy78=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_hashicorp_consul_sdk",
        importpath = "github.com/hashicorp/consul/sdk",
        sum = "h1:UOxjlb4xVNF93jak1mzzoBatyFju9nrkxpVwIp/QqxQ=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_hashicorp_errwrap",
        importpath = "github.com/hashicorp/errwrap",
        sum = "h1:OxrOeh75EUXMY8TBjag2fzXGZ40LB6IKw45YeGUDY2I=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_cleanhttp",
        importpath = "github.com/hashicorp/go-cleanhttp",
        sum = "h1:035FKYIWjmULyFRBKPs8TBQoi0x6d9G4xc9neXJWAZQ=",
        version = "v0.5.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_envparse",
        importpath = "github.com/hashicorp/go-envparse",
        sum = "h1:bE++6bhIsNCPLvgDZkYqo3nA+/PFI51pkrHdmPSDFPY=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_hclog",
        importpath = "github.com/hashicorp/go-hclog",
        sum = "h1:K4ev2ib4LdQETX5cSZBG0DVLk1jwGqSPXBjdah3veNs=",
        version = "v0.16.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_immutable_radix",
        importpath = "github.com/hashicorp/go-immutable-radix",
        sum = "h1:DKHmCUm2hRBK510BaiZlwvpD40f8bJFeZnpfm2KLowc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_msgpack",
        importpath = "github.com/hashicorp/go-msgpack",
        sum = "h1:zKjpN5BK/P5lMYrLmBHdBULWbJ0XpYR+7NGzqkZzoD4=",
        version = "v0.5.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_multierror",
        importpath = "github.com/hashicorp/go-multierror",
        sum = "h1:H5DkEtf6CXdFp0N0Em5UCwQpXMWke8IA0+lD48awMYo=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_net",
        importpath = "github.com/hashicorp/go.net",
        sum = "h1:sNCoNyDEvN1xa+X0baata4RdcpKwcMS6DH+xwfqPgjw=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_hashicorp_go_retryablehttp",
        importpath = "github.com/hashicorp/go-retryablehttp",
        sum = "h1:HJunrbHTDDbBb/ay4kxa1n+dLmttUlnP3V9oNE4hmsM=",
        version = "v0.6.6",
    )
    go_repository(
        name = "com_github_hashicorp_go_rootcerts",
        importpath = "github.com/hashicorp/go-rootcerts",
        sum = "h1:jzhAVGtqPKbwpyCPELlgNWhE1znq+qwJtW5Oi2viEzc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_sockaddr",
        importpath = "github.com/hashicorp/go-sockaddr",
        sum = "h1:ztczhD1jLxIRjVejw8gFomI1BQZOe2WoVOu0SyteCQc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_hashicorp_go_syslog",
        importpath = "github.com/hashicorp/go-syslog",
        sum = "h1:KaodqZuhUoZereWVIYmpUgZysurB1kBLX2j0MwMrUAE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_go_uuid",
        importpath = "github.com/hashicorp/go-uuid",
        sum = "h1:2gKiV6YVmrJ1i2CKKa9obLvRieoRGviZFL26PcT/Co8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_hashicorp_go_version",
        importpath = "github.com/hashicorp/go-version",
        sum = "h1:3vNe/fWF5CBgRIguda1meWhsZHy3m8gCJ5wx+dIzX/E=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_hashicorp_golang_lru",
        importpath = "github.com/hashicorp/golang-lru",
        sum = "h1:YDjusn29QI/Das2iO9M0BHnIbxPeyuCHsjMW+lJfyTc=",
        version = "v0.5.4",
    )
    go_repository(
        name = "com_github_hashicorp_hcl",
        importpath = "github.com/hashicorp/hcl",
        sum = "h1:0Anlzjpi4vEasTeNFn2mLJgTSwt0+6sfsiTG8qcWGx4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_logutils",
        importpath = "github.com/hashicorp/logutils",
        sum = "h1:dLEQVugN8vlakKOUE3ihGLTZJRB4j+M2cdTm/ORI65Y=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_mdns",
        importpath = "github.com/hashicorp/mdns",
        sum = "h1:WhIgCr5a7AaVH6jPUwjtRuuE7/RDufnUvzIr48smyxs=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_hashicorp_memberlist",
        importpath = "github.com/hashicorp/memberlist",
        sum = "h1:ouPxvwKYaNZe+eTcHxYP0EblPduVLvIPycul+vv8his=",
        version = "v0.1.6",
    )
    go_repository(
        name = "com_github_hashicorp_serf",
        importpath = "github.com/hashicorp/serf",
        sum = "h1:w2ZEHuK1297elT/WbZjUojVzpZA3BuPUusa9vdXXTjc=",
        version = "v0.8.6",
    )
    go_repository(
        name = "com_github_hashicorp_vault_api",
        importpath = "github.com/hashicorp/vault/api",
        sum = "h1:QcxC7FuqEl0sZaIjcXB/kNEeBa0DH5z57qbWBvZwLC4=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_hashicorp_vault_sdk",
        importpath = "github.com/hashicorp/vault/sdk",
        sum = "h1:e1ok06zGrWJW91rzRroyl5nRNqraaBe4d5hiKcVZuHM=",
        version = "v0.1.14-0.20200519221838-e0cfd64bc267",
    )
    go_repository(
        name = "com_github_hodgesds_perf_utils",
        importpath = "github.com/hodgesds/perf-utils",
        sum = "h1:7KlHGMuig4FRH5fNw68PV6xLmgTe7jKs9hgAcEAbioU=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_hpcloud_tail",
        importpath = "github.com/hpcloud/tail",
        sum = "h1:nfCOvKYfkgYP8hkirhJocXT2+zOD8yUNjXaWfTlyFKI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_huandu_xstrings",
        importpath = "github.com/huandu/xstrings",
        sum = "h1:yPeWdRnmynF7p+lLYz0H2tthW9lqhMJrQV/U7yy4wX0=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_hudl_fargo",
        importpath = "github.com/hudl/fargo",
        sum = "h1:0U6+BtN6LhaYuTnIJq4Wyq5cpn6O2kWrxAtcqBmYY6w=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_hugelgupf_socketpair",
        importpath = "github.com/hugelgupf/socketpair",
        sum = "h1:/jC7qQFrv8CrSJVmaolDVOxTfS9kc36uB6H40kdbQq8=",
        version = "v0.0.0-20190730060125-05d35a94e714",
    )
    go_repository(
        name = "com_github_iancoleman_strcase",
        importpath = "github.com/iancoleman/strcase",
        sum = "h1:05I4QRnGpI0m37iZQRuskXh+w77mr6Z41lwQzuHLwW0=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_ianlancetaylor_demangle",
        importpath = "github.com/ianlancetaylor/demangle",
        sum = "h1:rwmN+hgiyp8QyBqzdEX43lTjKAxaqCrYHaU5op5P9J8=",
        version = "v0.0.0-20220517205856-0058ec4f073c",
    )
    go_repository(
        name = "com_github_ibm_sarama",
        importpath = "github.com/IBM/sarama",
        sum = "h1:lL01NNg/iBeigUbT+wpPysuTYW6roHo6kc1QrffRf0k=",
        version = "v1.40.1",
    )
    go_repository(
        name = "com_github_illumos_go_kstat",
        importpath = "github.com/illumos/go-kstat",
        sum = "h1:hk4LPqXIY/c9XzRbe7dA6qQxaT6Axcbny0L/G5a4owQ=",
        version = "v0.0.0-20210513183136-173c9b0a9973",
    )
    go_repository(
        name = "com_github_imdario_mergo",
        importpath = "github.com/imdario/mergo",
        sum = "h1:lFzP57bqS/wsqKssCGmtLAb8A0wKjLGrve2q3PPVcBk=",
        version = "v0.3.13",
    )
    go_repository(
        name = "com_github_improbable_eng_grpc_web",
        importpath = "github.com/improbable-eng/grpc-web",
        sum = "h1:BN+7z6uNXZ1tQGcNAuaU1YjsLTApzkjt2tzCixLaUPQ=",
        version = "v0.15.0",
    )
    go_repository(
        name = "com_github_inconshreveable_mousetrap",
        importpath = "github.com/inconshreveable/mousetrap",
        sum = "h1:wN+x4NVGpMsO7ErUn/mUI3vEoE6Jt13X2s0bqwp9tc8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_influxdata_influxdb1_client",
        importpath = "github.com/influxdata/influxdb1-client",
        sum = "h1:/WZQPMZNsjZ7IlCpsLGdQBINg5bxKQ1K1sh6awxLtkA=",
        version = "v0.0.0-20191209144304-8bf82d3c094d",
    )
    go_repository(
        name = "com_github_infobloxopen_go_trees",
        importpath = "github.com/infobloxopen/go-trees",
        sum = "h1:w66aaP3c6SIQ0pi3QH1Tb4AMO3aWoEPxd1CNvLphbkA=",
        version = "v0.0.0-20200715205103-96a057b8dfb9",
    )
    go_repository(
        name = "com_github_insomniacslk_dhcp",
        importpath = "github.com/insomniacslk/dhcp",
        sum = "h1:PgxFEySCI41sH0mB7/2XswdXbUykQsRUGod8Rn+NubM=",
        version = "v0.0.0-20231016090811-6a2c8fbdcc1c",
    )
    go_repository(
        name = "com_github_intel_goresctrl",
        importpath = "github.com/intel/goresctrl",
        sum = "h1:K2D3GOzihV7xSBedGxONSlaw/un1LZgWsc9IfqipN4c=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_ishidawataru_sctp",
        importpath = "github.com/ishidawataru/sctp",
        sum = "h1:i2fYnDurfLlJH8AyyMOnkLHnHeP8Ff/DDpuZA/D3bPo=",
        version = "v0.0.0-20230406120618-7ff4192f6ff2",
    )
    go_repository(
        name = "com_github_j_keck_arping",
        importpath = "github.com/j-keck/arping",
        sum = "h1:hlLhuXgQkzIJTZuhMigvG/CuSkaspeaD9hRDk2zuiMI=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_jackc_chunkreader",
        importpath = "github.com/jackc/chunkreader",
        sum = "h1:4s39bBR8ByfqH+DKm8rQA3E1LHZWB9XWcrz8fqaZbe0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_chunkreader_v2",
        importpath = "github.com/jackc/chunkreader/v2",
        sum = "h1:i+RDz65UE+mmpjTfyz0MoVTnzeYxroil2G82ki7MGG8=",
        version = "v2.0.1",
    )
    go_repository(
        name = "com_github_jackc_pgconn",
        importpath = "github.com/jackc/pgconn",
        sum = "h1:smbxIaZA08n6YuxEX1sDyjV/qkbtUtkH20qLkR9MUR4=",
        version = "v1.14.1",
    )
    go_repository(
        name = "com_github_jackc_pgerrcode",
        importpath = "github.com/jackc/pgerrcode",
        sum = "h1:WAvSpGf7MsFuzAtK4Vk7R4EVe+liW4x83r4oWu0WHKw=",
        version = "v0.0.0-20201024163028-a0d42d470451",
    )
    go_repository(
        name = "com_github_jackc_pgio",
        importpath = "github.com/jackc/pgio",
        sum = "h1:g12B9UwVnzGhueNavwioyEEpAmqMe1E/BN9ES+8ovkE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgmock",
        importpath = "github.com/jackc/pgmock",
        sum = "h1:DadwsjnMwFjfWc9y5Wi/+Zz7xoE5ALHsRQlOctkOiHc=",
        version = "v0.0.0-20210724152146-4ad1a8207f65",
    )
    go_repository(
        name = "com_github_jackc_pgpassfile",
        importpath = "github.com/jackc/pgpassfile",
        sum = "h1:/6Hmqy13Ss2zCq62VdNG8tM1wchn8zjSGOBJ6icpsIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3",
        importpath = "github.com/jackc/pgproto3",
        sum = "h1:FYYE4yRw+AgI8wXIinMlNjBbp/UitDJwfj5LqqewP1A=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_jackc_pgproto3_v2",
        importpath = "github.com/jackc/pgproto3/v2",
        sum = "h1:7eY55bdBeCz1F2fTzSz69QC+pG46jYq9/jtSPiJ5nn0=",
        version = "v2.3.2",
    )
    go_repository(
        name = "com_github_jackc_pgservicefile",
        importpath = "github.com/jackc/pgservicefile",
        sum = "h1:bbPeKD0xmW/Y25WS6cokEszi5g+S0QxI/d45PkRi7Nk=",
        version = "v0.0.0-20221227161230-091c0ba34f0a",
    )
    go_repository(
        name = "com_github_jackc_pgtype",
        importpath = "github.com/jackc/pgtype",
        sum = "h1:y+xUdabmyMkJLyApYuPj38mW+aAIqCe5uuBB51rH3Vw=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_github_jackc_pgx_v4",
        importpath = "github.com/jackc/pgx/v4",
        sum = "h1:YP7G1KABtKpB5IHrO9vYwSrCOhs7p3uqhvhhQBptya0=",
        version = "v4.18.1",
    )
    go_repository(
        name = "com_github_jackc_pgx_v5",
        importpath = "github.com/jackc/pgx/v5",
        sum = "h1:cxFyXhxlvAifxnkKKdlxv8XqUf59tDlYjnV5YYfsJJY=",
        version = "v5.4.3",
    )
    go_repository(
        name = "com_github_jackc_puddle",
        importpath = "github.com/jackc/puddle",
        sum = "h1:gI8os0wpRXFd4FiAY2dWiqRK037tjj3t7rKFeO4X5iw=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_jackc_puddle_v2",
        importpath = "github.com/jackc/puddle/v2",
        sum = "h1:RhxXJtFG022u4ibrCSMSiu5aOq1i77R3OHKNJj77OAk=",
        version = "v2.2.1",
    )
    go_repository(
        name = "com_github_jarcoal_httpmock",
        importpath = "github.com/jarcoal/httpmock",
        sum = "h1:cHtVEcTxRSX4J0je7mWPfc9BpDpqzXSJ5HbymZmyHck=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_jbenet_go_context",
        importpath = "github.com/jbenet/go-context",
        sum = "h1:BQSFePA1RWJOlocH6Fxy8MmwDt+yVQYULKfN0RoTN8A=",
        version = "v0.0.0-20150711004518-d14ea06fba99",
    )
    go_repository(
        name = "com_github_jcmturner_aescts_v2",
        importpath = "github.com/jcmturner/aescts/v2",
        sum = "h1:9YKLH6ey7H4eDBXW8khjYslgyqG2xZikXP0EQFKrle8=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_dnsutils_v2",
        importpath = "github.com/jcmturner/dnsutils/v2",
        sum = "h1:lltnkeZGL0wILNvrNiVCR6Ro5PGU/SeBvVO/8c/iPbo=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_jcmturner_gofork",
        importpath = "github.com/jcmturner/gofork",
        sum = "h1:QH0l3hzAU1tfT3rZCnW5zXl+orbkNMMRGJfdJjHVETg=",
        version = "v1.7.6",
    )
    go_repository(
        name = "com_github_jcmturner_gokrb5_v8",
        importpath = "github.com/jcmturner/gokrb5/v8",
        sum = "h1:iTonLeSJOn7MVUtyMT+arAn5AKAPrkilzhGw8wE/Tq8=",
        version = "v8.4.3",
    )
    go_repository(
        name = "com_github_jcmturner_rpc_v2",
        importpath = "github.com/jcmturner/rpc/v2",
        sum = "h1:7FXXj8Ti1IaVFpSAziCZWNzbNuZmnvw/i6CqLNdWfZY=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_jeffashton_win_pdh",
        importpath = "github.com/JeffAshton/win_pdh",
        sum = "h1:UKkYhof1njT1/xq4SEg5z+VpTgjmNeHwPGRQl7takDI=",
        version = "v0.0.0-20161109143554-76bb4ee9f0ab",
    )
    go_repository(
        name = "com_github_jessevdk_go_flags",
        importpath = "github.com/jessevdk/go-flags",
        sum = "h1:4IU2WS7AumrZ/40jfhf4QVDMsQwqA7VEHozFRrGARJA=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_jgautheron_goconst",
        importpath = "github.com/jgautheron/goconst",
        sum = "h1:D5s1HIu80AcMGcqmk7fNIVptmAubVHHaj3v5Upex6Zs=",
        version = "v0.0.0-20170703170152-9740945f5dcb",
    )
    go_repository(
        name = "com_github_jhump_protoreflect",
        importpath = "github.com/jhump/protoreflect",
        sum = "h1:npqHz788dryJiR/l6K/RUQAyh2SwV91+d1dnh4RjO9w=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_jinzhu_gorm",
        importpath = "github.com/jinzhu/gorm",
        sum = "h1:HvrsqdhCW78xpJF67g1hMxS6eCToo9PZH4LDB8WKPac=",
        version = "v1.9.10",
    )
    go_repository(
        name = "com_github_jinzhu_inflection",
        importpath = "github.com/jinzhu/inflection",
        sum = "h1:K317FqzuhWc8YvSVlFMCCUb36O/S9MCKRDI7QkRKD/E=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jinzhu_now",
        importpath = "github.com/jinzhu/now",
        sum = "h1:/o9tlHleP7gOFmsnYNz3RGnqzefHA47wQpKrrdTIwXQ=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath",
        importpath = "github.com/jmespath/go-jmespath",
        sum = "h1:BEgLn5cpjn8UN1mAw4NjwDrS35OdebyEtFe+9YPoQUg=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_jmespath_go_jmespath_internal_testify",
        importpath = "github.com/jmespath/go-jmespath/internal/testify",
        sum = "h1:shLQSRRSCCPj3f2gpwzGwWFoC7ycTf1rcQZHOlsJ6N8=",
        version = "v1.5.1",
    )
    go_repository(
        name = "com_github_jmoiron_sqlx",
        importpath = "github.com/jmoiron/sqlx",
        sum = "h1:vFFPA71p1o5gAeqtEAwLU4dnX2napprKtHr7PYIcN3g=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_github_joefitzgerald_rainbow_reporter",
        importpath = "github.com/joefitzgerald/rainbow-reporter",
        sum = "h1:AuMG652zjdzI0YCCnXAqATtRBpGXMcAnrajcaTrSeuo=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_johncgriffin_overflow",
        importpath = "github.com/JohnCGriffin/overflow",
        sum = "h1:RGWPOewvKIROun94nF7v2cua9qP+thov/7M50KEoeSU=",
        version = "v0.0.0-20211019200055-46fa312c352c",
    )
    go_repository(
        name = "com_github_joho_godotenv",
        importpath = "github.com/joho/godotenv",
        sum = "h1:3l4+N6zfMWnkbPEXKng2o2/MR5mSwTrBih4ZEkkz1lg=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_jonboulle_clockwork",
        importpath = "github.com/jonboulle/clockwork",
        sum = "h1:UOGuzwb1PwsrDAObMuhUnj0p5ULPj8V/xJ7Kx9qUBdQ=",
        version = "v0.2.2",
    )
    go_repository(
        name = "com_github_josephspurrier_goversioninfo",
        importpath = "github.com/josephspurrier/goversioninfo",
        sum = "h1:Puhl12NSHUSALHSuzYwPYQkqa2E1+7SrtAPJorKK0C8=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_josharian_intern",
        importpath = "github.com/josharian/intern",
        sum = "h1:vlS4z54oSdjm0bgjRigI+G1HpF+tI+9rE5LLzOg8HmY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_josharian_native",
        importpath = "github.com/josharian/native",
        sum = "h1:uuaP0hAbW7Y4l0ZRQ6C9zfb7Mg1mbFKry/xzDAfmtLA=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_jpillora_backoff",
        importpath = "github.com/jpillora/backoff",
        sum = "h1:uvFg412JmmHBHw7iwprIxkPMI+sGQ4kzOWsMeHnm2EA=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_jsimonetti_rtnetlink",
        importpath = "github.com/jsimonetti/rtnetlink",
        replace = "github.com/jsimonetti/rtnetlink",
        sum = "h1:Z1BF0fRgcETPEa0Kt0MRk3yV5+kF1FWTni6KUFKrq2I=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_json_iterator_go",
        importpath = "github.com/json-iterator/go",
        sum = "h1:PV8peI4a0ysnczrg+LtxykD8LfKY9ML6u2jnxaEnrnM=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_jstemmer_go_junit_report",
        importpath = "github.com/jstemmer/go-junit-report",
        sum = "h1:6QPYqodiu3GuPL+7mfx+NwDdp2eTkp9IfEUpgAwUN0o=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_jtolds_gls",
        importpath = "github.com/jtolds/gls",
        sum = "h1:xdiiI2gbIgH/gLH7ADydsJ1uDOEzR8yvV7C0MuV77Wo=",
        version = "v4.20.0+incompatible",
    )
    go_repository(
        name = "com_github_juju_ratelimit",
        importpath = "github.com/juju/ratelimit",
        sum = "h1:+7AIFJVQ0EQgq/K9+0Krm7m530Du7tIz0METWzN0RgY=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_julienschmidt_httprouter",
        importpath = "github.com/julienschmidt/httprouter",
        sum = "h1:U0609e9tgbseu3rBINet9P48AI/D3oJs4dN7jwJOQ1U=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_jung_kurt_gofpdf",
        importpath = "github.com/jung-kurt/gofpdf",
        sum = "h1:PJr+ZMXIecYc1Ey2zucXdR73SMBtgjPgwa31099IMv0=",
        version = "v1.0.3-0.20190309125859-24315acbbda5",
    )
    go_repository(
        name = "com_github_k0kubun_colorstring",
        importpath = "github.com/k0kubun/colorstring",
        sum = "h1:uC1QfSlInpQF+M0ao65imhwqKnz3Q2z/d8PWZRMQvDM=",
        version = "v0.0.0-20150214042306-9440f1994b88",
    )
    go_repository(
        name = "com_github_k0kubun_pp",
        importpath = "github.com/k0kubun/pp",
        sum = "h1:EKhKbi34VQDWJtq+zpsKSEhkHHs9w2P8Izbq8IhLVSo=",
        version = "v2.3.0+incompatible",
    )
    go_repository(
        name = "com_github_kardianos_osext",
        importpath = "github.com/kardianos/osext",
        sum = "h1:iQTw/8FWTuc7uiaSepXwyf3o52HaUYcV+Tu66S3F5GA=",
        version = "v0.0.0-20190222173326-2bc1f35cddc0",
    )
    go_repository(
        name = "com_github_karrick_godirwalk",
        importpath = "github.com/karrick/godirwalk",
        sum = "h1:b4kY7nqDdioR/6qnbHQyDvmA17u5G1cZ6J+CZXwSWoI=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_github_kballard_go_shellquote",
        importpath = "github.com/kballard/go-shellquote",
        sum = "h1:Z9n2FFNUXsshfwJMBgNA0RU6/i7WVaAegv3PtuIHPMs=",
        version = "v0.0.0-20180428030007-95032a82bc51",
    )
    go_repository(
        name = "com_github_kevinburke_go_bindata",
        importpath = "github.com/kevinburke/go-bindata",
        sum = "h1:rqNOXZlqrYhMVVAsQx8wuc+LaA73YcfbQ407wAykyS8=",
        version = "v3.23.0+incompatible",
    )
    go_repository(
        name = "com_github_kevinburke_ssh_config",
        importpath = "github.com/kevinburke/ssh_config",
        sum = "h1:Coekwdh0v2wtGp9Gmz1Ze3eVRAWJMLokvN3QjdzCHLY=",
        version = "v0.0.0-20190725054713-01f96b0aa0cd",
    )
    go_repository(
        name = "com_github_kisielk_errcheck",
        importpath = "github.com/kisielk/errcheck",
        sum = "h1:e8esj/e4R+SAOwFwN+n3zr0nYeCyeweozKfO23MvHzY=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_kisielk_gotool",
        importpath = "github.com/kisielk/gotool",
        sum = "h1:AV2c/EiW3KqPNT9ZKl07ehoAGi4C5/01Cfbblndcapg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_klauspost_asmfmt",
        importpath = "github.com/klauspost/asmfmt",
        sum = "h1:4Ri7ox3EwapiOjCki+hw14RyKk201CN4rzyCJRFLpK4=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        sum = "h1:RlWWUY/Dr4fL8qk9YG7DTZ7PDgME2V4csBXA8L/ixi4=",
        version = "v1.17.2",
    )
    go_repository(
        name = "com_github_klauspost_cpuid_v2",
        importpath = "github.com/klauspost/cpuid/v2",
        sum = "h1:acbojRNwl3o09bUq+yDCtZFc1aiwaAAxtcn8YkZXnvk=",
        version = "v2.2.4",
    )
    go_repository(
        name = "com_github_knetic_govaluate",
        importpath = "github.com/Knetic/govaluate",
        sum = "h1:1G1pk05UrOh0NlF1oeaaix1x8XzrfjIDK47TY0Zehcw=",
        version = "v3.0.1-0.20171022003610-9aa49832a739+incompatible",
    )
    go_repository(
        name = "com_github_koneu_natend",
        importpath = "github.com/koneu/natend",
        sum = "h1:MFX8DxRnKMY/2M3H61iSsVbo/n3h0MWGmWNN1UViOU0=",
        version = "v0.0.0-20150829182554-ec0926ea948d",
    )
    go_repository(
        name = "com_github_konsorten_go_windows_terminal_sequences",
        importpath = "github.com/konsorten/go-windows-terminal-sequences",
        sum = "h1:CE8S1cTafDpPvMhIxNJKvHsGVBgn1xWYf1NbHQhywc8=",
        version = "v1.0.3",
    )
    go_repository(
        name = "com_github_kr_fs",
        importpath = "github.com/kr/fs",
        sum = "h1:Jskdu9ieNAYnjxsi0LbQp1ulIKZV1LAFgK1tWhpZgl8=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_kr_logfmt",
        importpath = "github.com/kr/logfmt",
        sum = "h1:T+h1c/A9Gawja4Y9mFVWj2vyii2bbUNDw3kt9VxK2EY=",
        version = "v0.0.0-20140226030751-b84e30acd515",
    )
    go_repository(
        name = "com_github_kr_pretty",
        importpath = "github.com/kr/pretty",
        sum = "h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_kr_pty",
        importpath = "github.com/kr/pty",
        sum = "h1:AkaSdXYQOWeaO3neb8EM634ahkXXe3jYbVh/F9lq+GI=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_kr_text",
        importpath = "github.com/kr/text",
        sum = "h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_ktrysmt_go_bitbucket",
        importpath = "github.com/ktrysmt/go-bitbucket",
        sum = "h1:C8dUGp0qkwncKtAnozHCbbqhptefzEd1I0sfnuy9rYQ=",
        version = "v0.6.4",
    )
    go_repository(
        name = "com_github_kylelemons_godebug",
        importpath = "github.com/kylelemons/godebug",
        sum = "h1:RPNrshWIDI6G2gRW9EHilWtl7Z6Sb1BR0xunSBf0SNc=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_labstack_echo",
        importpath = "github.com/labstack/echo",
        sum = "h1:pGRcYk231ExFAyoAjAfD85kQzRJCRI8bbnE7CX5OEgg=",
        version = "v3.3.10+incompatible",
    )
    go_repository(
        name = "com_github_labstack_echo_v4",
        importpath = "github.com/labstack/echo/v4",
        sum = "h1:wPOF1CE6gvt/kmbMR4dGzWvHMPT+sAEUJOwOTtvITVY=",
        version = "v4.9.0",
    )
    go_repository(
        name = "com_github_labstack_gommon",
        importpath = "github.com/labstack/gommon",
        sum = "h1:OomWaJXm7xR6L1HmEtGyQf26TEn7V6X88mktX9kee9o=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_leodido_go_urn",
        importpath = "github.com/leodido/go-urn",
        sum = "h1:XlAE/cm/ms7TE/VMVoduSpNBoyc2dOxHs5MZSwAN63Q=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_github_lestrrat_go_backoff_v2",
        importpath = "github.com/lestrrat-go/backoff/v2",
        sum = "h1:oNb5E5isby2kiro9AgdHLv5N5tint1AnDVVf2E2un5A=",
        version = "v2.0.8",
    )
    go_repository(
        name = "com_github_lestrrat_go_blackmagic",
        importpath = "github.com/lestrrat-go/blackmagic",
        sum = "h1:XzdxDbuQTz0RZZEmdU7cnQxUtFUzgCSPq8RCz4BxIi4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_lestrrat_go_httpcc",
        importpath = "github.com/lestrrat-go/httpcc",
        sum = "h1:ydWCStUeJLkpYyjLDHihupbn2tYmZ7m22BGkcvZZrIE=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_lestrrat_go_iter",
        importpath = "github.com/lestrrat-go/iter",
        sum = "h1:q8faalr2dY6o8bV45uwrxq12bRa1ezKrB6oM9FUgN4A=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_lestrrat_go_jwx",
        importpath = "github.com/lestrrat-go/jwx",
        sum = "h1:tAx93jN2SdPvFn08fHNAhqFJazn5mBBOB8Zli0g0otA=",
        version = "v1.2.25",
    )
    go_repository(
        name = "com_github_lestrrat_go_option",
        importpath = "github.com/lestrrat-go/option",
        sum = "h1:WqAWL8kh8VcSoD6xjSH34/1m8yxluXQbDeKNfvFeEO4=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_letsencrypt_pkcs11key_v4",
        importpath = "github.com/letsencrypt/pkcs11key/v4",
        sum = "h1:qLc/OznH7xMr5ARJgkZCCWk+EomQkiNTOoOF5LAgagc=",
        version = "v4.0.0",
    )
    go_repository(
        name = "com_github_lib_pq",
        importpath = "github.com/lib/pq",
        sum = "h1:YXG7RB+JIjhP29X+OtkiDnYaXQwpS4JEWq7dtCCRUEw=",
        version = "v1.10.9",
    )
    go_repository(
        name = "com_github_libopenstorage_openstorage",
        importpath = "github.com/libopenstorage/openstorage",
        sum = "h1:GLPam7/0mpdP8ZZtKjbfcXJBTIA/T1O6CBErVEFEyIM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_liggitt_tabwriter",
        importpath = "github.com/liggitt/tabwriter",
        sum = "h1:9TO3cAIGXtEhnIaL+V+BEER86oLrvS+kWobKpbJuye0=",
        version = "v0.0.0-20181228230101-89fcab3d43de",
    )
    go_repository(
        name = "com_github_lightstep_lightstep_tracer_common_golang_gogo",
        importpath = "github.com/lightstep/lightstep-tracer-common/golang/gogo",
        sum = "h1:143Bb8f8DuGWck/xpNUOckBVYfFbBTnLevfRZ1aVVqo=",
        version = "v0.0.0-20190605223551-bc2310a04743",
    )
    go_repository(
        name = "com_github_lightstep_lightstep_tracer_go",
        importpath = "github.com/lightstep/lightstep-tracer-go",
        sum = "h1:vi1F1IQ8N7hNWytK9DpJsUfQhGuNSc19z330K6vl4zk=",
        version = "v0.18.1",
    )
    go_repository(
        name = "com_github_linuxkit_virtsock",
        importpath = "github.com/linuxkit/virtsock",
        sum = "h1:jUp75lepDg0phMUJBCmvaeFDldD2N3S1lBuPwUTszio=",
        version = "v0.0.0-20201010232012-f8cee7dfc7a3",
    )
    go_repository(
        name = "com_github_lithammer_dedent",
        importpath = "github.com/lithammer/dedent",
        sum = "h1:VNzHMVCBNG1j0fh3OrsFRkVUwStdDArbgBWoPAffktY=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_lufia_iostat",
        importpath = "github.com/lufia/iostat",
        sum = "h1:tnCdZBIglgxD47RyD55kfWQcJMGzO+1QBziSQfesf2k=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star",
        importpath = "github.com/lyft/protoc-gen-star",
        sum = "h1:erE0rdztuaDq3bpGifD95wfoPrSZc95nGA6tbiNYh6M=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_star_v2",
        importpath = "github.com/lyft/protoc-gen-star/v2",
        sum = "h1:/3+/2sWyXeMLzKd1bX+ixWKgEMsULrIivpDsuaF441o=",
        version = "v2.0.3",
    )
    go_repository(
        name = "com_github_lyft_protoc_gen_validate",
        importpath = "github.com/lyft/protoc-gen-validate",
        sum = "h1:KNt/RhmQTOLr7Aj8PsJ7mTronaFyx80mRTT9qF261dA=",
        version = "v0.0.13",
    )
    go_repository(
        name = "com_github_magiconair_properties",
        importpath = "github.com/magiconair/properties",
        sum = "h1:ZC2Vc7/ZFkGmsVC9KvOjumD+G5lXy2RtTKyzRKO2BQ4=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_mailru_easyjson",
        importpath = "github.com/mailru/easyjson",
        sum = "h1:UGYAvKxe3sBsEDzO8ZeWOSlIQfWFlxbzLZe7hwFURr0=",
        version = "v0.7.7",
    )
    go_repository(
        name = "com_github_makenowjust_heredoc",
        importpath = "github.com/MakeNowJust/heredoc",
        sum = "h1:cXCdzVdstXyiTqTvfqk9SDHpKNjxuom+DOlyEeQ4pzQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_markbates_oncer",
        importpath = "github.com/markbates/oncer",
        sum = "h1:JgVTCPf0uBVcUSWpyXmGpgOc62nK5HWUBKAGc3Qqa5k=",
        version = "v0.0.0-20181203154359-bf2de49a0be2",
    )
    go_repository(
        name = "com_github_markbates_pkger",
        importpath = "github.com/markbates/pkger",
        sum = "h1:3MPelV53RnGSW07izx5xGxl4e/sdRD6zqseIk0rMASY=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_markbates_safe",
        importpath = "github.com/markbates/safe",
        sum = "h1:yjZkbvRM6IzKj9tlu/zMJLS0n/V351OZWRnF3QfaUxI=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_marstr_guid",
        importpath = "github.com/marstr/guid",
        sum = "h1:/M4H/1G4avsieL6BbUwCOBzulmoeKVP5ux/3mQNnbyI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_masterminds_goutils",
        importpath = "github.com/Masterminds/goutils",
        sum = "h1:zukEsf/1JZwCMgHiK3GZftabmxiCw4apj3a28RPBiVg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_masterminds_semver",
        importpath = "github.com/Masterminds/semver",
        sum = "h1:H65muMkzWKEuNDnfl9d70GUjFniHKHRbFPGBuZ3QEww=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_masterminds_semver_v3",
        importpath = "github.com/Masterminds/semver/v3",
        sum = "h1:hLg3sBzpNErnxhQtUy/mmLR2I9foDujNK030IGemrRc=",
        version = "v3.1.1",
    )
    go_repository(
        name = "com_github_masterminds_sprig",
        importpath = "github.com/Masterminds/sprig",
        sum = "h1:z4yfnGrZ7netVz+0EDJ0Wi+5VZCSYp4Z0m2dk6cEM60=",
        version = "v2.22.0+incompatible",
    )
    go_repository(
        name = "com_github_mattbaird_jsonpatch",
        importpath = "github.com/mattbaird/jsonpatch",
        sum = "h1:+J2gw7Bw77w/fbK7wnNJJDKmw1IbWft2Ul5BzrG1Qm8=",
        version = "v0.0.0-20171005235357-81af80346b1a",
    )
    go_repository(
        name = "com_github_mattn_go_colorable",
        importpath = "github.com/mattn/go-colorable",
        sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
        version = "v0.1.13",
    )
    go_repository(
        name = "com_github_mattn_go_ieproxy",
        importpath = "github.com/mattn/go-ieproxy",
        sum = "h1:qiyop7gCflfhwCzGyeT0gro3sF9AIg9HU98JORTkqfI=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_mattn_go_isatty",
        importpath = "github.com/mattn/go-isatty",
        sum = "h1:JITubQf0MOLdlGRuRq+jtsDlekdYPia9ZFsB8h/APPA=",
        version = "v0.0.19",
    )
    go_repository(
        name = "com_github_mattn_go_runewidth",
        importpath = "github.com/mattn/go-runewidth",
        sum = "h1:+xnbZSEeDbOIg5/mE6JF0w6n9duR1l3/WmbinWVwUuU=",
        version = "v0.0.14",
    )
    go_repository(
        name = "com_github_mattn_go_shellwords",
        importpath = "github.com/mattn/go-shellwords",
        sum = "h1:M2zGm7EW6UQJvDeQxo4T51eKPurbeFbe8WtebGE2xrk=",
        version = "v1.0.12",
    )
    go_repository(
        name = "com_github_mattn_go_sqlite3",
        importpath = "github.com/mattn/go-sqlite3",
        sum = "h1:mCRHCLDUBXgpKAqIKsaAaAsrAlbkeomtRFKXh2L6YIM=",
        version = "v1.14.17",
    )
    go_repository(
        name = "com_github_mattn_go_xmlrpc",
        importpath = "github.com/mattn/go-xmlrpc",
        sum = "h1:Y6WEMLEsqs3RviBrAa1/7qmbGB7DVD3brZIbqMbQdGY=",
        version = "v0.0.3",
    )
    go_repository(
        name = "com_github_mattn_go_zglob",
        importpath = "github.com/mattn/go-zglob",
        sum = "h1:xsEx/XUoVlI6yXjqBK062zYhRTZltCNmYPx6v+8DNaY=",
        version = "v0.0.1",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        sum = "h1:mmDVorXM7PCGKw94cs5zkfA9PSy5pEvNWRP0ET0TIVo=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extensions_v2",
        importpath = "github.com/matttproud/golang_protobuf_extensions/v2",
        sum = "h1:jWpvCLoY8Z/e3VKvlsiIGKtc+UG6U5vzxaoagmhXfyg=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_maxbrunsfeld_counterfeiter_v6",
        importpath = "github.com/maxbrunsfeld/counterfeiter/v6",
        sum = "h1:g+4J5sZg6osfvEfkRZxJ1em0VT95/UOZgi/l7zi1/oE=",
        version = "v6.2.2",
    )
    go_repository(
        name = "com_github_mdlayher_arp",
        importpath = "github.com/mdlayher/arp",
        sum = "h1:ql8x//rJsHMjS+qqEag8n3i4azw1QneKh5PieH9UEbY=",
        version = "v0.0.0-20220512170110-6706a2966875",
    )
    go_repository(
        name = "com_github_mdlayher_ethernet",
        importpath = "github.com/mdlayher/ethernet",
        sum = "h1:2oDp6OOhLxQ9JBoUuysVz9UZ9uI6oLUbvAZu0x8o+vE=",
        version = "v0.0.0-20220221185849-529eae5b6118",
    )
    go_repository(
        name = "com_github_mdlayher_ethtool",
        importpath = "github.com/mdlayher/ethtool",
        replace = "github.com/monogon-dev/ethtool",
        sum = "h1:Y1CNfkifx2lVkPUX7KC4Hn9qIPUoSm3RiWPwEwy6VA8=",
        version = "v0.0.0-20231122193313-e9c21a3a83cb",
    )
    go_repository(
        name = "com_github_mdlayher_genetlink",
        importpath = "github.com/mdlayher/genetlink",
        sum = "h1:KdrNKe+CTu+IbZnm/GVUMXSqBBLqcGpRDa0xkQy56gw=",
        version = "v1.3.2",
    )
    go_repository(
        name = "com_github_mdlayher_kobject",
        importpath = "github.com/mdlayher/kobject",
        sum = "h1:JmrZTpS0GAyMV4ZQVVH/AS0Y6r2PbnYNSRUuRX+HOLA=",
        version = "v0.0.0-20200520190114-19ca17470d7d",
    )
    go_repository(
        name = "com_github_mdlayher_netlink",
        importpath = "github.com/mdlayher/netlink",
        sum = "h1:/UtM3ofJap7Vl4QWCPDGXY8d3GIY2UGSDbK+QWmY8/g=",
        version = "v1.7.2",
    )
    go_repository(
        name = "com_github_mdlayher_packet",
        importpath = "github.com/mdlayher/packet",
        sum = "h1:3Up1NG6LZrsgDVn6X4L9Ge/iyRyxFEFD9o6Pr3Q1nQY=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_mdlayher_socket",
        importpath = "github.com/mdlayher/socket",
        sum = "h1:ilICZmJcQz70vrWVes1MFera4jGiWNocSkykwwoy3XI=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_mdlayher_wifi",
        importpath = "github.com/mdlayher/wifi",
        sum = "h1:y8wYRUXwok5CtUZOXT3egghYesX0O79E3ALl+SIDm9Q=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_mgutz_ansi",
        importpath = "github.com/mgutz/ansi",
        sum = "h1:j7+1HpAFS1zy5+Q4qx1fWh90gTKwiN4QCGoY9TWyyO4=",
        version = "v0.0.0-20170206155736-9520e82c474b",
    )
    go_repository(
        name = "com_github_mibk_dupl",
        importpath = "github.com/mibk/dupl",
        sum = "h1:aZc3jqrF9n0tUHwHt/+jsRxA8cRgA0Gdl56M7W7PoqE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_microsoft_go_mssqldb",
        importpath = "github.com/microsoft/go-mssqldb",
        sum = "h1:p2rpHIL7TlSv1QrbXJUAcbyRKnIT0C9rRkH2E4OjLn8=",
        version = "v0.21.0",
    )
    go_repository(
        name = "com_github_microsoft_go_winio",
        importpath = "github.com/Microsoft/go-winio",
        sum = "h1:9/kr64B9VUZrLm5YYwbGtUJnMgqWVOdUAXu6Migciow=",
        version = "v0.6.1",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim",
        build_file_proto_mode = "disable",
        importpath = "github.com/Microsoft/hcsshim",
        sum = "h1:68vKo2VN8DE9AdN4tnkWnmdhqdbpUFM8OF3Airm7fz8=",
        version = "v0.11.4",
    )
    go_repository(
        name = "com_github_microsoft_hcsshim_test",
        importpath = "github.com/Microsoft/hcsshim/test",
        sum = "h1:4FA+QBaydEHlwxg0lMN3rhwoDaQy6LKhVWR4qvq4BuA=",
        version = "v0.0.0-20210227013316-43a75bb4edd3",
    )
    go_repository(
        name = "com_github_miekg_dns",
        importpath = "github.com/miekg/dns",
        sum = "h1:GoQ4hpsj0nFLYe+bWiCToyrBEJXkQfOOIvFGFy0lEgo=",
        version = "v1.1.55",
    )
    go_repository(
        name = "com_github_miekg_pkcs11",
        importpath = "github.com/miekg/pkcs11",
        sum = "h1:Ugu9pdy6vAYku5DEpVWVFPYnzV+bxB+iRdbuFSu7TvU=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_mikioh_ipaddr",
        importpath = "github.com/mikioh/ipaddr",
        sum = "h1:RlZweED6sbSArvlE924+mUcZuXKLBHA35U7LN621Bws=",
        version = "v0.0.0-20190404000644-d465c8ab6721",
    )
    go_repository(
        name = "com_github_minio_asm2plan9s",
        importpath = "github.com/minio/asm2plan9s",
        sum = "h1:AMFGa4R4MiIpspGNG7Z948v4n35fFGB3RR3G/ry4FWs=",
        version = "v0.0.0-20200509001527-cdd76441f9d8",
    )
    go_repository(
        name = "com_github_minio_c2goasm",
        importpath = "github.com/minio/c2goasm",
        sum = "h1:+n/aFZefKZp7spd8DFdX7uMikMLXX4oubIzJF4kv/wI=",
        version = "v0.0.0-20190812172519-36a3d3bbc4f3",
    )
    go_repository(
        name = "com_github_minio_sha256_simd",
        importpath = "github.com/minio/sha256-simd",
        sum = "h1:v1ta+49hkWZyvaKwrQB8elexRqm6Y0aMLjCNsrYxo6g=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mistifyio_go_zfs",
        importpath = "github.com/mistifyio/go-zfs",
        sum = "h1:aKW/4cBs+yK6gpqU3K/oIwk9Q/XICqd3zOX/UFuvqmk=",
        version = "v2.1.2-0.20190413222219-f784269be439+incompatible",
    )
    go_repository(
        name = "com_github_mistifyio_go_zfs_v3",
        importpath = "github.com/mistifyio/go-zfs/v3",
        sum = "h1:YaoXgBePoMA12+S1u/ddkv+QqxcfiZK4prI6HPnkFiU=",
        version = "v3.0.1",
    )
    go_repository(
        name = "com_github_mitchellh_cli",
        importpath = "github.com/mitchellh/cli",
        sum = "h1:iGBIsUe3+HZ/AD/Vd7DErOt5sU9fa8Uj7A2s1aggv1Y=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_copystructure",
        importpath = "github.com/mitchellh/copystructure",
        sum = "h1:Laisrj+bAB6b/yJwB5Bt3ITZhGJdqmxquMKeZ+mmkFQ=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_homedir",
        importpath = "github.com/mitchellh/go-homedir",
        sum = "h1:lukF9ziXFxDFPkA1vsr5zpc1XuPDn/wFntq5mG+4E0Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_testing_interface",
        importpath = "github.com/mitchellh/go-testing-interface",
        sum = "h1:fzU/JVNcaqHQEcVFAKeR41fkiLdIPrefOvVG1VZ96U0=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_go_wordwrap",
        importpath = "github.com/mitchellh/go-wordwrap",
        sum = "h1:TLuKupo69TCn6TQSyGxwI1EblZZEsQ0vMlAFQflz0v0=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_mitchellh_gox",
        importpath = "github.com/mitchellh/gox",
        sum = "h1:lfGJxY7ToLJQjHHwi0EX6uYBdK78egf954SQl13PQJc=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_mitchellh_iochan",
        importpath = "github.com/mitchellh/iochan",
        sum = "h1:C+X3KsSTLFVBr/tK1eYN/vs4rJcvsiLU338UhYPJWeY=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mitchellh_mapstructure",
        importpath = "github.com/mitchellh/mapstructure",
        sum = "h1:6h7AQ0yhTcIsmFmnAwQls75jp2Gzs4iB8W7pjMO+rqo=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_mitchellh_osext",
        importpath = "github.com/mitchellh/osext",
        sum = "h1:2+myh5ml7lgEU/51gbeLHfKGNfgEQQIWrlbdaOsidbQ=",
        version = "v0.0.0-20151018003038-5e2d6d41470f",
    )
    go_repository(
        name = "com_github_mitchellh_reflectwalk",
        importpath = "github.com/mitchellh/reflectwalk",
        sum = "h1:FVzMWA5RllMAKIdUSC8mdWo3XtwoecrH79BY70sEEpE=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_moby_ipvs",
        importpath = "github.com/moby/ipvs",
        sum = "h1:ONN4pGaZQgAx+1Scz5RvWV4Q7Gb+mvfRh3NsPS+1XQQ=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_moby_locker",
        importpath = "github.com/moby/locker",
        sum = "h1:fOXqR41zeveg4fFODix+1Ch4mj/gT0NE1XJbp/epuBg=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_moby_spdystream",
        importpath = "github.com/moby/spdystream",
        sum = "h1:cjW1zVyyoiM0T7b6UoySUFqzXMoqRckQtXwGPiBhOM8=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_moby_sys_mountinfo",
        importpath = "github.com/moby/sys/mountinfo",
        sum = "h1:BzJjoreD5BMFNmD9Rus6gdd1pLuecOFPt8wC+Vygl78=",
        version = "v0.6.2",
    )
    go_repository(
        name = "com_github_moby_sys_sequential",
        importpath = "github.com/moby/sys/sequential",
        sum = "h1:OPvI35Lzn9K04PBbCLW0g4LcFAJgHsvXsRyewg5lXtc=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_moby_sys_signal",
        importpath = "github.com/moby/sys/signal",
        sum = "h1:25RW3d5TnQEoKvRbEKUGay6DCQ46IxAVTT9CUMgmsSI=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_moby_sys_symlink",
        importpath = "github.com/moby/sys/symlink",
        sum = "h1:tk1rOM+Ljp0nFmfOIBtlV3rTDlWOwFRhjEeAhZB0nZc=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_moby_sys_user",
        importpath = "github.com/moby/sys/user",
        sum = "h1:WmZ93f5Ux6het5iituh9x2zAG7NFY9Aqi49jjE1PaQg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_moby_term",
        importpath = "github.com/moby/term",
        sum = "h1:HfkjXDfhgVaN5rmueG8cL8KKeFNecRCXFhaJ2qZ5SKA=",
        version = "v0.0.0-20221205130635-1aeaba878587",
    )
    go_repository(
        name = "com_github_modern_go_concurrent",
        importpath = "github.com/modern-go/concurrent",
        sum = "h1:TRLaZ9cD/w8PVh93nsPXa1VrQ6jlwL5oN8l14QlcNfg=",
        version = "v0.0.0-20180306012644-bacd9c7ef1dd",
    )
    go_repository(
        name = "com_github_modern_go_reflect2",
        importpath = "github.com/modern-go/reflect2",
        sum = "h1:xBagoLtFs94CBntxluKeaWgTMpvLxC4ur3nMaC9Gz0M=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_modocache_gover",
        importpath = "github.com/modocache/gover",
        sum = "h1:8Q0qkMVC/MmWkpIdlvZgcv2o2jrlF6zqVOh7W5YHdMA=",
        version = "v0.0.0-20171022184752-b58185e213c5",
    )
    go_repository(
        name = "com_github_mohae_deepcopy",
        importpath = "github.com/mohae/deepcopy",
        sum = "h1:RWengNIwukTxcDr9M+97sNutRR1RKhG96O6jWumTTnw=",
        version = "v0.0.0-20170929034955-c48cc78d4826",
    )
    go_repository(
        name = "com_github_monochromegane_go_gitignore",
        importpath = "github.com/monochromegane/go-gitignore",
        sum = "h1:n6/2gBQ3RWajuToeY6ZtZTIKv2v7ThUy5KKusIT0yc0=",
        version = "v0.0.0-20200626010858-205db1a8cc00",
    )
    go_repository(
        name = "com_github_montanaflynn_stats",
        importpath = "github.com/montanaflynn/stats",
        sum = "h1:iruDEfMl2E6fbMZ9s0scYfZQ84/6SPL6zC8ACM2oIL0=",
        version = "v0.0.0-20171201202039-1bf9dbcd8cbe",
    )
    go_repository(
        name = "com_github_morikuni_aec",
        importpath = "github.com/morikuni/aec",
        sum = "h1:nP9CBfwrvYnBRgY6qfDQkygYDmYwOilePFkwzv4dU8A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_mrunalp_fileutils",
        importpath = "github.com/mrunalp/fileutils",
        sum = "h1:F+S7ZlNKnrwHfSwdlgNSkKo67ReVf8o9fel6C3dkm/Q=",
        version = "v0.5.1",
    )
    go_repository(
        name = "com_github_muesli_reflow",
        importpath = "github.com/muesli/reflow",
        sum = "h1:99aDTygRy9yEwggATz+ZLrDFRsjRog5BqbAfsr47Ztw=",
        version = "v0.0.0-20191128061954-86f094cbed14",
    )
    go_repository(
        name = "com_github_munnerz_goautoneg",
        importpath = "github.com/munnerz/goautoneg",
        sum = "h1:C3w9PqII01/Oq1c1nUAm88MOHcQC9l5mIlSMApZMrHA=",
        version = "v0.0.0-20191010083416-a7dc8b61c822",
    )
    go_repository(
        name = "com_github_mutecomm_go_sqlcipher_v4",
        importpath = "github.com/mutecomm/go-sqlcipher/v4",
        sum = "h1:sV1tWCWGAVlPhNGT95Q+z/txFxuhAYWwHD1afF5bMZg=",
        version = "v4.4.0",
    )
    go_repository(
        name = "com_github_mwitkow_go_conntrack",
        importpath = "github.com/mwitkow/go-conntrack",
        sum = "h1:KUppIJq7/+SVif2QVs3tOP0zanoHgBEVAwHxUSIzRqU=",
        version = "v0.0.0-20190716064945-2f068394615f",
    )
    go_repository(
        name = "com_github_mwitkow_go_proto_validators",
        importpath = "github.com/mwitkow/go-proto-validators",
        replace = "github.com/mwitkow/go-proto-validators",
        sum = "h1:qRlmpTzm2pstMKKzTdvwPCF5QfBNURSlAgN/R+qbKos=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_mwitkow_grpc_proxy",
        importpath = "github.com/mwitkow/grpc-proxy",
        sum = "h1:0xuRacu/Zr+jX+KyLLPPktbwXqyOvnOPUQmMLzX1jxU=",
        version = "v0.0.0-20181017164139-0f1106ef9c76",
    )
    go_repository(
        name = "com_github_mxk_go_flowrate",
        importpath = "github.com/mxk/go-flowrate",
        sum = "h1:y5//uYreIhSUg3J1GEMiLbxo1LJaP8RfCpH6pymGZus=",
        version = "v0.0.0-20140419014527-cca7078d478f",
    )
    go_repository(
        name = "com_github_nakagami_firebirdsql",
        importpath = "github.com/nakagami/firebirdsql",
        sum = "h1:P48LjvUQpTReR3TQRbxSeSBsMXzfK0uol7eRcr7VBYQ=",
        version = "v0.0.0-20190310045651-3c02a58cfed8",
    )
    go_repository(
        name = "com_github_nats_io_jwt",
        importpath = "github.com/nats-io/jwt",
        sum = "h1:+RB5hMpXUUA2dfxuhBTEkMOrYmM+gKIZYS1KjSostMI=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_nats_io_nats_go",
        importpath = "github.com/nats-io/nats.go",
        sum = "h1:ik3HbLhZ0YABLto7iX80pZLPw/6dx3T+++MZJwLnMrQ=",
        version = "v1.9.1",
    )
    go_repository(
        name = "com_github_nats_io_nats_server_v2",
        importpath = "github.com/nats-io/nats-server/v2",
        sum = "h1:i2Ly0B+1+rzNZHHWtD4ZwKi+OU5l+uQo1iDHZ2PmiIc=",
        version = "v2.1.2",
    )
    go_repository(
        name = "com_github_nats_io_nkeys",
        importpath = "github.com/nats-io/nkeys",
        sum = "h1:6JrEfig+HzTH85yxzhSVbjHRJv9cn0p6n3IngIcM5/k=",
        version = "v0.1.3",
    )
    go_repository(
        name = "com_github_nats_io_nuid",
        importpath = "github.com/nats-io/nuid",
        sum = "h1:5iA8DT8V7q8WK2EScv2padNa/rTESc1KdnPw4TC2paw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_github_ncw_swift",
        importpath = "github.com/ncw/swift",
        sum = "h1:4DQRPj35Y41WogBxyhOXlrI37nzGlyEcsforeudyYPQ=",
        version = "v1.0.47",
    )
    go_repository(
        name = "com_github_neo4j_neo4j_go_driver",
        importpath = "github.com/neo4j/neo4j-go-driver",
        sum = "h1:fhFP5RliM2HW/8XdcO5QngSfFli9GcRIpMXvypTQt6E=",
        version = "v1.8.1-0.20200803113522-b626aa943eba",
    )
    go_repository(
        name = "com_github_networkplumbing_go_nft",
        importpath = "github.com/networkplumbing/go-nft",
        sum = "h1:eKapmyVUt/3VGfhYaDos5yeprm+LPt881UeksmKKZHY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_niemeyer_pretty",
        importpath = "github.com/niemeyer/pretty",
        sum = "h1:fD57ERR4JtEqsWbfPhv4DMiApHyliiK5xCTNVSPiaAs=",
        version = "v0.0.0-20200227124842-a10e7caefd8e",
    )
    go_repository(
        name = "com_github_nishanths_predeclared",
        importpath = "github.com/nishanths/predeclared",
        sum = "h1:3f0nxAmdj/VoCGN/ijdMy7bj6SBagaqYg1B0hu8clMA=",
        version = "v0.0.0-20200524104333-86fad755b4d3",
    )
    go_repository(
        name = "com_github_nxadm_tail",
        importpath = "github.com/nxadm/tail",
        sum = "h1:nPr65rt6Y5JFSKQO7qToXr7pePgD6Gwiw05lkbyAQTE=",
        version = "v1.4.8",
    )
    go_repository(
        name = "com_github_nytimes_gziphandler",
        importpath = "github.com/NYTimes/gziphandler",
        sum = "h1:ZUDjpQae29j0ryrS0u/B8HZfJBtBQHjqw2rQ2cqUQ3I=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_oklog_oklog",
        importpath = "github.com/oklog/oklog",
        sum = "h1:wVfs8F+in6nTBMkA7CbRw+zZMIB7nNM825cM1wuzoTk=",
        version = "v0.3.2",
    )
    go_repository(
        name = "com_github_oklog_run",
        importpath = "github.com/oklog/run",
        sum = "h1:Ru7dDtJNOyC66gQ5dQmaCa0qIsAUFY3sFpK1Xk8igrw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_oklog_ulid",
        importpath = "github.com/oklog/ulid",
        sum = "h1:EGfNDEx6MqHz8B3uNV6QAib1UR2Lm97sHi3ocA6ESJ4=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_olekukonko_tablewriter",
        importpath = "github.com/olekukonko/tablewriter",
        sum = "h1:P2Ga83D34wi1o9J6Wh1mRuqd4mF/x/lgBS7N7AbDhec=",
        version = "v0.0.5",
    )
    go_repository(
        name = "com_github_oneofone_xxhash",
        importpath = "github.com/OneOfOne/xxhash",
        sum = "h1:31czK/TI9sNkxIKfaUfGlU47BAxQ0ztGgd9vPyqimf8=",
        version = "v1.2.8",
    )
    go_repository(
        name = "com_github_onsi_ginkgo",
        importpath = "github.com/onsi/ginkgo",
        sum = "h1:8xi0RTUf59SOSfEtZMvwTvXYMzG4gV23XVHOZiXNtnE=",
        version = "v1.16.5",
    )
    go_repository(
        name = "com_github_onsi_ginkgo_v2",
        importpath = "github.com/onsi/ginkgo/v2",
        sum = "h1:79HwNRBAZHOEwrczrgSOPy+eFTTlIGELKy5as+ClttY=",
        version = "v2.15.0",
    )
    go_repository(
        name = "com_github_onsi_gomega",
        importpath = "github.com/onsi/gomega",
        sum = "h1:54UJxxj6cPInHS3a35wm6BK/F9nHYueZ1NVujHDrnXE=",
        version = "v1.31.0",
    )
    go_repository(
        name = "com_github_op_go_logging",
        importpath = "github.com/op/go-logging",
        sum = "h1:lDH9UUVJtmYCjyT0CI4q8xvlXPxeZ0gYCVvWbmPlp88=",
        version = "v0.0.0-20160315200505-970db520ece7",
    )
    go_repository(
        name = "com_github_open_policy_agent_opa",
        importpath = "github.com/open-policy-agent/opa",
        sum = "h1:qocVAKyjrqMjCqsU02S/gHyLr4AQQ9xMtuV1kKnnyhM=",
        version = "v0.42.2",
    )
    go_repository(
        name = "com_github_opencontainers_go_digest",
        importpath = "github.com/opencontainers/go-digest",
        sum = "h1:apOUWs51W5PlhuyGyz9FCeeBIOUDA/6nW8Oi/yOhh5U=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opencontainers_image_spec",
        importpath = "github.com/opencontainers/image-spec",
        sum = "h1:YWuSjZCQAPM8UUBLkYUk1e+rZcvWHJmFb6i6rM44Xs8=",
        version = "v1.1.0-rc2.0.20221005185240-3a7f492d3f1b",
    )
    go_repository(
        name = "com_github_opencontainers_runc",
        build_tags = ["seccomp"],  #keep
        importpath = "github.com/opencontainers/runc",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:runc-add-cdeps.patch",
        ],
        sum = "h1:BOIssBaW1La0/qbNZHXOOa71dZfZEQOzW7dqQf3phss=",
        version = "v1.1.12",
    )
    go_repository(
        name = "com_github_opencontainers_runtime_spec",
        importpath = "github.com/opencontainers/runtime-spec",
        sum = "h1:HHUyrt9mwHUjtasSbXSMvs4cyFxh+Bll4AjJ9odEGpg=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_opencontainers_runtime_tools",
        importpath = "github.com/opencontainers/runtime-tools",
        sum = "h1:DmNGcqH3WDbV5k8OJ+esPWbqUOX5rMLR2PMvziDMJi0=",
        version = "v0.9.1-0.20221107090550-2e043c6bd626",
    )
    go_repository(
        name = "com_github_opencontainers_selinux",
        importpath = "github.com/opencontainers/selinux",
        sum = "h1:+5Zbo97w3Lbmb3PeqQtpmTkMwsW5nRI3YaLpt7tQ7oU=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_opentracing_basictracer_go",
        importpath = "github.com/opentracing/basictracer-go",
        sum = "h1:YyUAhaEfjoWXclZVJ9sGoNct7j4TVk7lZWlQw5UXuoo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_opentracing_contrib_go_observer",
        importpath = "github.com/opentracing-contrib/go-observer",
        sum = "h1:lM6RxxfUMrYL/f8bWEUqdXrANWtrL7Nndbm9iFN0DlU=",
        version = "v0.0.0-20170622124052-a52f23424492",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        importpath = "github.com/opentracing/opentracing-go",
        sum = "h1:uEJPy/1a5RIPAJ0Ov+OIO8OxWu77jEv+1B0VhjKrZUs=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_openzipkin_contrib_zipkin_go_opentracing",
        importpath = "github.com/openzipkin-contrib/zipkin-go-opentracing",
        sum = "h1:uhcF5Jd7rP9DVEL10Siffyepr6SvlKbUsjH5JpNCRi8=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_openzipkin_zipkin_go",
        importpath = "github.com/openzipkin/zipkin-go",
        sum = "h1:zjqfqHjUpPmB3c1GlCvvgsM1G4LkvqQbBDueDOCg/jA=",
        version = "v0.4.2",
    )
    go_repository(
        name = "com_github_oschwald_geoip2_golang",
        importpath = "github.com/oschwald/geoip2-golang",
        sum = "h1:uvD3O6fXAXs+usU+UGExshpdP13GAqp4GBrzN7IgKZc=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_github_oschwald_maxminddb_golang",
        importpath = "github.com/oschwald/maxminddb-golang",
        sum = "h1:aSXMqYR/EPNjGE8epgqwDay+P30hCBZIveY0WZbAWh0=",
        version = "v1.11.0",
    )
    go_repository(
        name = "com_github_otiai10_copy",
        importpath = "github.com/otiai10/copy",
        sum = "h1:HvG945u96iNadPoG2/Ja2+AUJeW5YuFQMixq9yirC+k=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_otiai10_curr",
        importpath = "github.com/otiai10/curr",
        sum = "h1:TJIWdbX0B+kpNagQrjgq8bCMrbhiuX73M2XwgtDMoOI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_otiai10_mint",
        importpath = "github.com/otiai10/mint",
        sum = "h1:BCmzIS3n71sGfHB5NMNDB3lHYPz8fWSkCAErHed//qc=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_outcaste_io_ristretto",
        importpath = "github.com/outcaste-io/ristretto",
        sum = "h1:KCItuNIGJZcursqHr3ghO7fc5ddZLEHspL9UR0cQM64=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_packethost_packngo",
        importpath = "github.com/packethost/packngo",
        replace = "github.com/monogon-dev/packngo",
        sum = "h1:sxSnvzB4iDBNhUBqXME/ETqjF4vX0mURE85T/I/Mr0o=",
        version = "v0.0.0-20240122175436-ecbd9eb00ddb",
    )
    go_repository(
        name = "com_github_pact_foundation_pact_go",
        importpath = "github.com/pact-foundation/pact-go",
        sum = "h1:OYkFijGHoZAYbOIb1LWXrwKQbMMRUv1oQ89blD2Mh2Q=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_github_pascaldekloe_goe",
        importpath = "github.com/pascaldekloe/goe",
        sum = "h1:Lgl0gzECD8GnQ5QCWA8o6BtfL6mDH5rQgM4/fX3avOs=",
        version = "v0.0.0-20180627143212-57f6aae5913c",
    )
    go_repository(
        name = "com_github_pborman_uuid",
        importpath = "github.com/pborman/uuid",
        sum = "h1:J7Q5mO4ysT1dv8hyrUGHb9+ooztCXu1D8MY8DZYsu3g=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_pelletier_go_buffruneio",
        importpath = "github.com/pelletier/go-buffruneio",
        sum = "h1:U4t4R6YkofJ5xHm3dJzuRpPZ0mr5MMCoAWooScCR7aA=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_pelletier_go_toml",
        importpath = "github.com/pelletier/go-toml",
        sum = "h1:4yBQzkHv+7BHq2PQUZF3Mx0IYxG7LsP222s7Agd3ve8=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_github_pelletier_go_toml_v2",
        importpath = "github.com/pelletier/go-toml/v2",
        sum = "h1:0ctb6s9mE31h0/lhu+J6OPmVeDxJn+kYnJc2jZR9tGQ=",
        version = "v2.0.8",
    )
    go_repository(
        name = "com_github_performancecopilot_speed",
        importpath = "github.com/performancecopilot/speed",
        sum = "h1:2WnRzIquHa5QxaJKShDkLM+sc0JPuwhXzK8OYOyt3Vg=",
        version = "v3.0.0+incompatible",
    )
    go_repository(
        name = "com_github_peterbourgon_diskv",
        importpath = "github.com/peterbourgon/diskv",
        sum = "h1:UBdAOUP5p4RWqPBg048CAvpKN+vxiaj6gdUUzhl4XmI=",
        version = "v2.0.1+incompatible",
    )
    go_repository(
        name = "com_github_pganalyze_pg_query_go_v4",
        importpath = "github.com/pganalyze/pg_query_go/v4",
        patch_args = ["-p1"],
        patches = [
            # pganalyze depends on libpg_query, and this patch adds a cdep to
            # it to the generated BUILD files.
            "//third_party/go/patches:pganalyze-add-cdeps.patch",
        ],
        sum = "h1:cNLqyiVMasV7YGWyYV+fkXyHp32gDfXVNCqoHztEGNk=",
        version = "v4.2.3",
    )
    go_repository(
        name = "com_github_philhofer_fwd",
        importpath = "github.com/philhofer/fwd",
        sum = "h1:bnDivRJ1EWPjUIRXV5KfORO897HTbpFAQddBdE8t7Gw=",
        version = "v1.1.2",
    )
    go_repository(
        name = "com_github_phpdave11_gofpdf",
        importpath = "github.com/phpdave11/gofpdf",
        sum = "h1:KPKiIbfwbvC/wOncwhrpRdXVj2CZTCFlw4wnoyjtHfQ=",
        version = "v1.4.2",
    )
    go_repository(
        name = "com_github_phpdave11_gofpdi",
        importpath = "github.com/phpdave11/gofpdi",
        sum = "h1:o61duiW8M9sMlkVXWlvP92sZJtGKENvW3VExs6dZukQ=",
        version = "v1.0.13",
    )
    go_repository(
        name = "com_github_pierrec_lz4",
        importpath = "github.com/pierrec/lz4",
        sum = "h1:9UY3+iC23yxF0UfGaYrGplQ+79Rg+h/q9FV9ix19jjM=",
        version = "v2.6.1+incompatible",
    )
    go_repository(
        name = "com_github_pierrec_lz4_v4",
        importpath = "github.com/pierrec/lz4/v4",
        sum = "h1:kV4Ip+/hUBC+8T6+2EgburRtkE9ef4nbY3f4dFhGjMc=",
        version = "v4.1.17",
    )
    go_repository(
        name = "com_github_pingcap_errors",
        importpath = "github.com/pingcap/errors",
        sum = "h1:+FZIDR/D97YOPik4N4lPDaUcLDF/EQPogxtlHB2ZZRM=",
        version = "v0.11.5-0.20210425183316-da1aaba5fb63",
    )
    go_repository(
        name = "com_github_pingcap_failpoint",
        importpath = "github.com/pingcap/failpoint",
        sum = "h1:CgbKAHto5CQgWM9fSBIvaxsJHuGP0uM74HXtv3MyyGQ=",
        version = "v0.0.0-20220801062533-2eaa32854a6c",
    )
    go_repository(
        name = "com_github_pingcap_log",
        importpath = "github.com/pingcap/log",
        sum = "h1:ELiPxACz7vdo1qAvvaWJg1NrYFoY6gqAh/+Uo6aXdD8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_pingcap_tidb_parser",
        build_file_generation = "on",
        importpath = "github.com/pingcap/tidb/parser",
        sum = "h1:4HUBkIZs+4j+tbXGm5/B0yjB66OTz218HDKA6VrhO7U=",
        version = "v0.0.0-20231010133155-38cb4f3312be",
    )
    go_repository(
        name = "com_github_pkg_browser",
        importpath = "github.com/pkg/browser",
        sum = "h1:KoWmjvw+nsYOo29YJK9vDA65RGE3NrOnUtO7a+RF9HU=",
        version = "v0.0.0-20210911075715-681adbf594b8",
    )
    go_repository(
        name = "com_github_pkg_diff",
        importpath = "github.com/pkg/diff",
        sum = "h1:aoZm08cpOy4WuID//EZDgcC4zIxODThtZNPirFr42+A=",
        version = "v0.0.0-20210226163009-20ebb0f2a09e",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
    go_repository(
        name = "com_github_pkg_profile",
        importpath = "github.com/pkg/profile",
        sum = "h1:F++O52m40owAmADcojzM+9gyjmMOY/T4oYJkgFDH8RE=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_pkg_sftp",
        importpath = "github.com/pkg/sftp",
        sum = "h1:I2qBYMChEhIjOgazfJmV3/mZM256btk6wkCDRmW7JYs=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_github_pkg_xattr",
        importpath = "github.com/pkg/xattr",
        sum = "h1:dhclzL6EqOXNaPDWqoeb9tIxATfBSmjqL0b4DpSjwRw=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_pmezard_go_difflib",
        importpath = "github.com/pmezard/go-difflib",
        sum = "h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_posener_complete",
        importpath = "github.com/posener/complete",
        sum = "h1:ccV59UEOTzVDnDUEFdT95ZzHVZ+5+158q8+SJb2QV5w=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_pquerna_cachecontrol",
        importpath = "github.com/pquerna/cachecontrol",
        sum = "h1:yJMy84ti9h/+OEWa752kBTKv4XC30OtVVHYv/8cTqKc=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        sum = "h1:rl2sfwZMtSthVU752MqfjQozy7blglC+1SOtjMAMh+Q=",
        version = "v1.17.0",
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        sum = "h1:VQw1hfvPvk3Uv6Qf29VrPF32JB6rtbgI6cYPYQjL0Qw=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        sum = "h1:2BGz0eBc2hdMDLnO/8n0jeB3oPrt2D08CekT0lneoxM=",
        version = "v0.45.0",
    )
    go_repository(
        name = "com_github_prometheus_community_go_runit",
        importpath = "github.com/prometheus-community/go-runit",
        sum = "h1:uTWEj/Fn2RoLdfg/etSqwzgYNOYPrARx1BHUN052tGA=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_prometheus_exporter_toolkit",
        importpath = "github.com/prometheus/exporter-toolkit",
        sum = "h1:yOAzZTi4M22ZzVxD+fhy1URTuNRj/36uQJJ5S8IPza8=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_prometheus_node_exporter",
        build_tags = ["nobtrfs", "nozfs", "notapestats", "norapl"],  # keep
        importpath = "github.com/prometheus/node_exporter",
        sum = "h1:7MVpSdfWrThNo0SlldhUyAVFZ7LWbC9+QJRzB4QmkE8=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        sum = "h1:jluTpSng7V9hY0O2R9DzzJHYb2xULk9VTR1V1R/k6Bo=",
        version = "v0.12.0",
    )
    go_repository(
        name = "com_github_prometheus_tsdb",
        importpath = "github.com/prometheus/tsdb",
        sum = "h1:YZcsG11NqnK4czYLrWd9mpEuAJIHVQLwdrleYfszMAA=",
        version = "v0.7.1",
    )
    go_repository(
        name = "com_github_pseudomuto_protoc_gen_doc",
        build_directives = [
            "gazelle:resolve go github.com/mwitkow/go-proto-validators @com_github_mwitkow_go_proto_validators//:validators_golang",
        ],
        importpath = "github.com/pseudomuto/protoc-gen-doc",
        patch_args = ["-p1"],
        pre_patches = [
            "//third_party/go/patches:protoc-gen-doc-no-gogo.patch",
        ],
        sum = "h1:pHZp0MEiT68jrZV8js8BS7E9ZEnlSLegoQbbtXj5lfo=",
        version = "v1.5.0",
    )
    go_repository(
        name = "com_github_pseudomuto_protokit",
        build_file_proto_mode = "disable",
        importpath = "github.com/pseudomuto/protokit",
        sum = "h1:hlnBDcy3YEDXH7kc9gV+NLaN0cDzhDvD1s7Y6FZ8RpM=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_quic_go_qpack",
        importpath = "github.com/quic-go/qpack",
        sum = "h1:Cr9BXA1sQS2SmDUWjSofMPNKmvF6IiIfDRmgU0w1ZCo=",
        version = "v0.4.0",
    )
    go_repository(
        name = "com_github_quic_go_qtls_go1_20",
        importpath = "github.com/quic-go/qtls-go1-20",
        sum = "h1:O4BLOM3hwfVF3AcktIylQXyl7Yi2iBNVy5QsV+ySxbg=",
        version = "v0.3.1",
    )
    go_repository(
        name = "com_github_quic_go_quic_go",
        importpath = "github.com/quic-go/quic-go",
        sum = "h1:ke8B73yMCWGq9MfrCCAw0Uzdm7GaViC3i39dsIdDlH4=",
        version = "v0.37.4",
    )
    go_repository(
        name = "com_github_rabbitmq_amqp091_go",
        importpath = "github.com/rabbitmq/amqp091-go",
        sum = "h1:RejT1SBUim5doqcL6s7iN6SBmsQqyTgXb1xMlH0h1hA=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_github_rcrowley_go_metrics",
        importpath = "github.com/rcrowley/go-metrics",
        sum = "h1:N/ElC8H3+5XpJzTSTfLsJV/mx9Q9g7kxmchpfZyxgzM=",
        version = "v0.0.0-20201227073835-cf1acfcdf475",
    )
    go_repository(
        name = "com_github_redis_go_redis_v9",
        importpath = "github.com/redis/go-redis/v9",
        sum = "h1:r2ctp2J2+TcXTVIyPU6++FniED/Nyo4SDMKvLtpszx0=",
        version = "v9.0.0",
    )
    go_repository(
        name = "com_github_remyoudompheng_bigfft",
        importpath = "github.com/remyoudompheng/bigfft",
        sum = "h1:W09IVJc94icq4NjY3clb7Lk8O1qJ8BdBEF8z0ibU0rE=",
        version = "v0.0.0-20230129092748-24d4a6f8daec",
    )
    go_repository(
        name = "com_github_richardartoul_molecule",
        importpath = "github.com/richardartoul/molecule",
        sum = "h1:Qp27Idfgi6ACvFQat5+VJvlYToylpM/hcyLBI3WaKPA=",
        version = "v1.0.1-0.20221107223329-32cfee06a052",
    )
    go_repository(
        name = "com_github_rivo_uniseg",
        importpath = "github.com/rivo/uniseg",
        sum = "h1:S1pD9weZBuJdFmowNwbpi7BJ8TNftyUImj/0WQi72jY=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_riza_io_grpc_go",
        importpath = "github.com/riza-io/grpc-go",
        sum = "h1:2HxQKFVE7VuYstcJ8zqpN84VnAoJ4dCL6YFhJewNcHQ=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_rmohr_bazeldnf",
        importpath = "github.com/rmohr/bazeldnf",
        sum = "h1:xYSQoQHuCZY+2mZJtt+2KN0G6TmeEOPR6cxPZomTXX4=",
        version = "v0.5.4",
    )
    go_repository(
        name = "com_github_robfig_cron_v3",
        importpath = "github.com/robfig/cron/v3",
        sum = "h1:WdRxkvbJztn8LMz/QEvLN5sBU+xKpSqwwUO1Pjr4qDs=",
        version = "v3.0.1",
    )
    go_repository(
        name = "com_github_rogpeppe_fastuuid",
        importpath = "github.com/rogpeppe/fastuuid",
        sum = "h1:Ppwyp6VYCF1nvBTXL3trRso7mXMlRrw9ooo375wvi2s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_rogpeppe_go_internal",
        importpath = "github.com/rogpeppe/go-internal",
        sum = "h1:TMyTOH3F/DB16zRVcYyreMH6GnZZrwQVAoYjRBZyWFQ=",
        version = "v1.10.0",
    )
    go_repository(
        name = "com_github_rs_cors",
        importpath = "github.com/rs/cors",
        sum = "h1:P2KMzcFwrPoSjkF1WLRPsp3UMLyql8L4v9hQpVeK5so=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_github_rs_xid",
        importpath = "github.com/rs/xid",
        sum = "h1:mhH9Nq+C1fY2l1XIpgxIiUOfNpRBYH1kKcr+qfKgjRc=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_rs_zerolog",
        importpath = "github.com/rs/zerolog",
        sum = "h1:uPRuwkWF4J6fGsJ2R0Gn2jB1EQiav9k3S6CSdygQJXY=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_github_russross_blackfriday",
        importpath = "github.com/russross/blackfriday",
        sum = "h1:KqfZb0pUVN2lYqZUYRddxF4OR8ZMURnJIG5Y3VRLtww=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_github_russross_blackfriday_v2",
        importpath = "github.com/russross/blackfriday/v2",
        sum = "h1:JIOH55/0cWyOuilr9/qlrm0BSXldqnqwMsf35Ld67mk=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_ruudk_golang_pdf417",
        importpath = "github.com/ruudk/golang-pdf417",
        sum = "h1:K1Xf3bKttbF+koVGaX5xngRIZ5bVjbmPnaxE/dR08uY=",
        version = "v0.0.0-20201230142125-a7e3863a1245",
    )
    go_repository(
        name = "com_github_ryanuber_columnize",
        importpath = "github.com/ryanuber/columnize",
        sum = "h1:UFr9zpz4xgTnIE5yIMtWAMngCdZ9p/+q6lTbgelo80M=",
        version = "v0.0.0-20160712163229-9b3edd62028f",
    )
    go_repository(
        name = "com_github_ryanuber_go_glob",
        importpath = "github.com/ryanuber/go-glob",
        sum = "h1:iQh3xXAumdQ+4Ufa5b25cRpC5TYKlno6hsv6Cb3pkBk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_safchain_ethtool",
        importpath = "github.com/safchain/ethtool",
        sum = "h1:gimQJpsI6sc1yIqP/y8GYgiXn/NjgvpM0RNoWLVVmP0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_github_samuel_go_zookeeper",
        importpath = "github.com/samuel/go-zookeeper",
        sum = "h1:p3Vo3i64TCLY7gIfzeQaUJ+kppEO5WQG3cL8iE8tGHU=",
        version = "v0.0.0-20190923202752-2cc03de413da",
    )
    go_repository(
        name = "com_github_sassoftware_go_rpmutils",
        importpath = "github.com/sassoftware/go-rpmutils",
        replace = "github.com/rmohr/go-rpmutils",
        sum = "h1:IPm+rJWi7P9mkoFQjloDEGThxlM71dzDrC2DZSjgnL8=",
        version = "v0.1.2-0.20201215123907-5acf7436c00d",
    )
    go_repository(
        name = "com_github_satori_go_uuid",
        importpath = "github.com/satori/go.uuid",
        sum = "h1:0uYX9dsZ2yD7q2RtLRtPSdGDWzjeM3TbMJP9utgA0ww=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_savsgio_dictpool",
        importpath = "github.com/savsgio/dictpool",
        sum = "h1:rmMl4fXJhKMNWl+K+r/fq4FbbKI+Ia2m9hYBLm2h4G4=",
        version = "v0.0.0-20221023140959-7bf2e61cea94",
    )
    go_repository(
        name = "com_github_savsgio_gotils",
        importpath = "github.com/savsgio/gotils",
        sum = "h1:8Iv5m6xEo1NR1AvpV+7XmhI4r39LGNzwUL4YpMuL5vk=",
        version = "v0.0.0-20230208104028-c358bd845dee",
    )
    go_repository(
        name = "com_github_sbezverk_nfproxy",
        importpath = "github.com/sbezverk/nfproxy",
        patch_args = ["-p1"],
        pre_patches = [
            "//third_party/go/patches:nfproxy-adapt-to-k8s-1.28.patch",
        ],
        sum = "h1:nlKDckTP5mXfdpUST7l4cU6fGzMlkv/CVg+EVOsClS8=",
        version = "v0.0.0-20210112155058-0d98b4a69f0c",
    )
    go_repository(
        name = "com_github_sbezverk_nftableslib",
        importpath = "github.com/sbezverk/nftableslib",
        sum = "h1:Zi8wDJAUmUYqWX6A9ma84xOhFvwR7O10K7pvISM+SvM=",
        version = "v0.0.0-20210111145735-b08b2d804e1f",
    )
    go_repository(
        name = "com_github_sbezverk_nftableslib_e2e_setenv",
        importpath = "github.com/sbezverk/nftableslib/e2e/setenv",
        sum = "h1:/8S2Bo/8GZhXJBEf9yywFO2NlVzwH1jNdu4UIk48SCs=",
        version = "v0.0.0-20191010164456-029e0d78cdb1",
    )
    go_repository(
        name = "com_github_sclevine_agouti",
        importpath = "github.com/sclevine/agouti",
        sum = "h1:8IBJS6PWz3uTlMP3YBIR5f+KAldcGuOeFkFbUWfBgK4=",
        version = "v3.0.0+incompatible",
    )
    go_repository(
        name = "com_github_sclevine_spec",
        importpath = "github.com/sclevine/spec",
        sum = "h1:1Jwdf9jSfDl9NVmt8ndHqbTZ7XCCPbh1jI3hkDBHVYA=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_sean_seed",
        importpath = "github.com/sean-/seed",
        sum = "h1:nn5Wsu0esKSJiIVhscUtVbo7ada43DJhG55ua/hjS5I=",
        version = "v0.0.0-20170313163322-e2103e2c3529",
    )
    go_repository(
        name = "com_github_seccomp_libseccomp_golang",
        importpath = "github.com/seccomp/libseccomp-golang",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:libseccomp.patch",
        ],
        sum = "h1:aA4bp+/Zzi0BnWZ2F1wgNBs5gTpm+na2rWM6M9YjLpY=",
        version = "v0.10.0",
    )
    go_repository(
        name = "com_github_secure_systems_lab_go_securesystemslib",
        importpath = "github.com/secure-systems-lab/go-securesystemslib",
        sum = "h1:OwvJ5jQf9LnIAS83waAjPbcMsODrTQUpJ02eNLUoxBg=",
        version = "v0.7.0",
    )
    go_repository(
        name = "com_github_segmentio_kafka_go",
        importpath = "github.com/segmentio/kafka-go",
        sum = "h1:4ujULpikzHG0HqKhjumDghFjy/0RRCSl/7lbriwQAH0=",
        version = "v0.4.29",
    )
    go_repository(
        name = "com_github_sergi_go_diff",
        importpath = "github.com/sergi/go-diff",
        sum = "h1:XU+rvMAioB0UC3q1MFrIQy4Vo5/4VsRDQQXHsEya6xQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_shopify_logrus_bugsnag",
        importpath = "github.com/Shopify/logrus-bugsnag",
        sum = "h1:UrqY+r/OJnIp5u0s1SbQ8dVfLCZJsnvazdBP5hS4iRs=",
        version = "v0.0.0-20171204204709-577dee27f20d",
    )
    go_repository(
        name = "com_github_shopify_sarama",
        importpath = "github.com/Shopify/sarama",
        sum = "h1:rtiODsvY4jW6nUV6n3K+0gx/8WlAwVt+Ixt6RIvpYyo=",
        version = "v1.22.0",
    )
    go_repository(
        name = "com_github_shopify_toxiproxy",
        importpath = "github.com/Shopify/toxiproxy",
        sum = "h1:TKdv8HiTLgE5wdJuEML90aBgNWsokNbMijUGhmcoBJc=",
        version = "v2.1.4+incompatible",
    )
    go_repository(
        name = "com_github_shopspring_decimal",
        importpath = "github.com/shopspring/decimal",
        sum = "h1:2Usl1nmF/WZucqkFZhnfFYxxxu8LG21F6nPQBE5gKV8=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_github_shurcool_sanitized_anchor_name",
        importpath = "github.com/shurcooL/sanitized_anchor_name",
        sum = "h1:PdmoCO6wvbs+7yrJyMORt4/BmY5IYyJwS/kOiWx8mHo=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_siebenmann_go_kstat",
        importpath = "github.com/siebenmann/go-kstat",
        sum = "h1:GfSdC6wKfTGcgCS7BtzF5694Amne1pGCSTY252WhlEY=",
        version = "v0.0.0-20210513183136-173c9b0a9973",
    )
    go_repository(
        name = "com_github_sirupsen_logrus",
        importpath = "github.com/sirupsen/logrus",
        sum = "h1:dueUQJ1C2q9oE3F7wvmSGAaVtTmUizReu6fjN8uqzbQ=",
        version = "v1.9.3",
    )
    go_repository(
        name = "com_github_smartystreets_assertions",
        importpath = "github.com/smartystreets/assertions",
        sum = "h1:UVQPSSmc3qtTi+zPPkCXvZX9VvW/xT/NsRvKfwY81a8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_smartystreets_go_aws_auth",
        importpath = "github.com/smartystreets/go-aws-auth",
        sum = "h1:hp2CYQUINdZMHdvTdXtPOY2ainKl4IoMcpAXEf2xj3Q=",
        version = "v0.0.0-20180515143844-0c1422d1fdb9",
    )
    go_repository(
        name = "com_github_smartystreets_goconvey",
        importpath = "github.com/smartystreets/goconvey",
        sum = "h1:fv0U8FUIMPNf1L9lnHLvLhgicrIVChEkdzIKYqbNC9s=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_github_smartystreets_gunit",
        importpath = "github.com/smartystreets/gunit",
        sum = "h1:RyPDUFcJbvtXlhJPk7v+wnxZRY2EUokhEYl2EJOPToI=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_snowflakedb_gosnowflake",
        importpath = "github.com/snowflakedb/gosnowflake",
        sum = "h1:EJDdDi74YbYt1ty164ge3fMZ0eVZ6KA7b1zmAa/wnRo=",
        version = "v1.6.3",
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        importpath = "github.com/soheilhy/cmux",
        sum = "h1:jjzc5WVemNEDTLwv9tlmemhC73tI08BNOIGwBOo10Js=",
        version = "v0.1.5",
    )
    go_repository(
        name = "com_github_sony_gobreaker",
        importpath = "github.com/sony/gobreaker",
        sum = "h1:oMnRNZXX5j85zso6xCPRNPtmAycat+WcoKbklScLDgQ=",
        version = "v0.4.1",
    )
    go_repository(
        name = "com_github_spaolacci_murmur3",
        importpath = "github.com/spaolacci/murmur3",
        sum = "h1:qLC7fQah7D6K1B0ujays3HV9gkFtllcxhzImRR7ArPQ=",
        version = "v0.0.0-20180118202830-f09979ecbc72",
    )
    go_repository(
        name = "com_github_spf13_afero",
        importpath = "github.com/spf13/afero",
        sum = "h1:j49Hj62F0n+DaZ1dDCvhABaPNSGNkt32oRFxI33IEMw=",
        version = "v1.9.2",
    )
    go_repository(
        name = "com_github_spf13_cast",
        importpath = "github.com/spf13/cast",
        sum = "h1:oget//CVOEoFewqQxwr0Ej5yjygnqGkvggSE/gB35Q8=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_spf13_cobra",
        importpath = "github.com/spf13/cobra",
        sum = "h1:hyqWnYt1ZQShIddO5kBpj3vu05/++x6tJ6dg8EC572I=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_spf13_jwalterweatherman",
        importpath = "github.com/spf13/jwalterweatherman",
        sum = "h1:XHEdyB+EcvlqZamSM4ZOMGlc93t6AcsBEu9Gc1vn7yk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_spf13_pflag",
        importpath = "github.com/spf13/pflag",
        sum = "h1:iy+VFUOCP1a+8yFto/drg2CJ5u0yRoB7fZw3DKv/JXA=",
        version = "v1.0.5",
    )
    go_repository(
        name = "com_github_spf13_viper",
        importpath = "github.com/spf13/viper",
        sum = "h1:xVKxvI7ouOI5I+U9s2eeiUfMaWBVoXA3AWskkrqK0VM=",
        version = "v1.7.0",
    )
    go_repository(
        name = "com_github_sqlc_dev_sqlc",
        # Hard exclude wasm library. For some reason, this is required even
        # though we use the nowasm build tag (adding it here didn't help
        # either).
        build_extra_args = [
            "-exclude=internal/ext/wasm/wasm.go",
        ],
        # sqlc ships with pregenerated protobuf stubs.
        build_file_proto_mode = "disable",
        importpath = "github.com/sqlc-dev/sqlc",
        sum = "h1:h3hnieAGeaqcdKK2zqmy2fHi7S3EkQRR/Gc3qhKuEMg=",
        version = "v1.23.0",
    )
    go_repository(
        name = "com_github_src_d_gcfg",
        importpath = "github.com/src-d/gcfg",
        sum = "h1:xXbNR5AlLSA315x2UO+fTSSAXCDf+Ar38/6oyGbDKQ4=",
        version = "v1.4.0",
    )
    go_repository(
        name = "com_github_stefanberger_go_pkcs11uri",
        importpath = "github.com/stefanberger/go-pkcs11uri",
        sum = "h1:lIOOHPEbXzO3vnmx2gok1Tfs31Q8GQqKLc8vVqyQq/I=",
        version = "v0.0.0-20201008174630-78d3cae3a980",
    )
    go_repository(
        name = "com_github_stoewer_go_strcase",
        importpath = "github.com/stoewer/go-strcase",
        sum = "h1:g0eASXYtp+yvN9fK8sH94oCIk0fau9uV1/ZdJ0AVEzs=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_streadway_amqp",
        importpath = "github.com/streadway/amqp",
        sum = "h1:WhxRHzgeVGETMlmVfqhRn8RIeeNoPr2Czh33I4Zdccw=",
        version = "v0.0.0-20190827072141-edfb9018d271",
    )
    go_repository(
        name = "com_github_streadway_handy",
        importpath = "github.com/streadway/handy",
        sum = "h1:AhmOdSHeswKHBjhsLs/7+1voOxT+LLrSk/Nxvk35fug=",
        version = "v0.0.0-20190108123426-d5acb3125c2a",
    )
    go_repository(
        name = "com_github_stretchr_objx",
        importpath = "github.com/stretchr/objx",
        sum = "h1:1zr/of2m5FGMsad5YfcqgdqdWrIhu+EBEJRhR1U7z/c=",
        version = "v0.5.0",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        sum = "h1:CcVxjf3Q8PM0mHUKJCdn+eZZtm5yQwehR5yeSVQQcUk=",
        version = "v1.8.4",
    )
    go_repository(
        name = "com_github_stripe_safesql",
        importpath = "github.com/stripe/safesql",
        sum = "h1:xiefmCDd8c35PVSGrL2FhBiaKxviXnGziBDOpOejeBE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "com_github_subosito_gotenv",
        importpath = "github.com/subosito/gotenv",
        sum = "h1:Slr1R9HxAlEKefgq5jn9U+DnETlIUa6HfgEzj0g5d7s=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_syndtr_gocapability",
        importpath = "github.com/syndtr/gocapability",
        sum = "h1:kdXcSzyDtseVEc4yCz2qF8ZrQvIDBJLl4S1c3GCXmoI=",
        version = "v0.0.0-20200815063812-42c35b437635",
    )
    go_repository(
        name = "com_github_syndtr_goleveldb",
        importpath = "github.com/syndtr/goleveldb",
        sum = "h1:vfofYNRScrDdvS342BElfbETmL1Aiz3i2t0zfRj16Hs=",
        version = "v1.0.1-0.20220721030215-126854af5e6d",
    )
    go_repository(
        name = "com_github_tchap_go_patricia",
        importpath = "github.com/tchap/go-patricia",
        sum = "h1:JvoDL7JSoIP2HDE8AbDH3zC8QBPxmzYe32HHy5yQ+Ck=",
        version = "v2.2.6+incompatible",
    )
    go_repository(
        name = "com_github_tchap_go_patricia_v2",
        importpath = "github.com/tchap/go-patricia/v2",
        sum = "h1:6rQp39lgIYZ+MHmdEq4xzuk1t7OdC35z/xm0BGhTkes=",
        version = "v2.3.1",
    )
    go_repository(
        name = "com_github_tidwall_btree",
        importpath = "github.com/tidwall/btree",
        sum = "h1:5P+9WU8ui5uhmcg3SoPyTwoI0mVyZ1nps7YQzTZFkYM=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_tidwall_buntdb",
        importpath = "github.com/tidwall/buntdb",
        sum = "h1:8KOzf5Gg97DoCMSOgcwZjnM0FfROtq0fcZkPW54oGKU=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_tidwall_gjson",
        importpath = "github.com/tidwall/gjson",
        sum = "h1:ikuZsLdhr8Ws0IdROXUS1Gi4v9Z4pGqpX/CvJkxvfpo=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_github_tidwall_grect",
        importpath = "github.com/tidwall/grect",
        sum = "h1:dA3oIgNgWdSspFzn1kS4S/RDpZFLrIxAZOdJKjYapOg=",
        version = "v0.1.4",
    )
    go_repository(
        name = "com_github_tidwall_match",
        importpath = "github.com/tidwall/match",
        sum = "h1:+Ho715JplO36QYgwN9PGYNhgZvoUSc9X2c80KVTi+GA=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_github_tidwall_pretty",
        importpath = "github.com/tidwall/pretty",
        sum = "h1:RWIZEg2iJ8/g6fDDYzMpobmaoGh5OLl4AXtGUGPcqCs=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_tidwall_rtred",
        importpath = "github.com/tidwall/rtred",
        sum = "h1:exmoQtOLvDoO8ud++6LwVsAMTu0KPzLTUrMln8u1yu8=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_tidwall_tinyqueue",
        importpath = "github.com/tidwall/tinyqueue",
        sum = "h1:SpNEvEggbpyN5DIReaJ2/1ndroY8iyEGxPYxoSaymYE=",
        version = "v0.1.1",
    )
    go_repository(
        name = "com_github_tinylib_msgp",
        importpath = "github.com/tinylib/msgp",
        sum = "h1:FCXC1xanKO4I8plpHGH2P7koL/RzZs12l/+r7vakfm0=",
        version = "v1.1.8",
    )
    go_repository(
        name = "com_github_tj_assert",
        importpath = "github.com/tj/assert",
        sum = "h1:Rw8kxzWo1mr6FSaYXjQELRe88y2KdfynXdnK72rdjtA=",
        version = "v0.0.0-20171129193455-018094318fb0",
    )
    go_repository(
        name = "com_github_tj_go_elastic",
        importpath = "github.com/tj/go-elastic",
        sum = "h1:eGaGNxrtoZf/mBURsnNQKDR7u50Klgcf2eFDQEnc8Bc=",
        version = "v0.0.0-20171221160941-36157cbbebc2",
    )
    go_repository(
        name = "com_github_tj_go_kinesis",
        importpath = "github.com/tj/go-kinesis",
        sum = "h1:m74UWYy+HBs+jMFR9mdZU6shPewugMyH5+GV6LNgW8w=",
        version = "v0.0.0-20171128231115-08b17f58cb1b",
    )
    go_repository(
        name = "com_github_tj_go_spin",
        importpath = "github.com/tj/go-spin",
        sum = "h1:lhdWZsvImxvZ3q1C5OIB7d72DuOwP4O2NdBg9PyzNds=",
        version = "v1.1.0",
    )
    go_repository(
        name = "com_github_tmc_grpc_websocket_proxy",
        importpath = "github.com/tmc/grpc-websocket-proxy",
        sum = "h1:6fotK7otjonDflCTK0BCfls4SPy3NcCVb5dqqmbRknE=",
        version = "v0.0.0-20220101234140-673ab2c3ae75",
    )
    go_repository(
        name = "com_github_tmthrgd_go_hex",
        importpath = "github.com/tmthrgd/go-hex",
        sum = "h1:9lRDQMhESg+zvGYmW5DyG0UqvY96Bu5QYsTLvCHdrgo=",
        version = "v0.0.0-20190904060850-447a3041c3bc",
    )
    go_repository(
        name = "com_github_tomasen_realip",
        importpath = "github.com/tomasen/realip",
        sum = "h1:fb190+cK2Xz/dvi9Hv8eCYJYvIGUTN2/KLq1pT6CjEc=",
        version = "v0.0.0-20180522021738-f0c99a92ddce",
    )
    go_repository(
        name = "com_github_tsenart_deadcode",
        importpath = "github.com/tsenart/deadcode",
        sum = "h1:vY5WqiEon0ZSTGM3ayVVi+twaHKHDFUVloaQ/wug9/c=",
        version = "v0.0.0-20160724212837-210d2dc333e9",
    )
    go_repository(
        name = "com_github_tv42_httpunix",
        importpath = "github.com/tv42/httpunix",
        sum = "h1:u6SKchux2yDvFQnDHS3lPnIRmfVJ5Sxy3ao2SIdysLQ=",
        version = "v0.0.0-20191220191345-2ba4b9c3382c",
    )
    go_repository(
        name = "com_github_twitchtv_twirp",
        importpath = "github.com/twitchtv/twirp",
        sum = "h1:s5WnVKMhC4Xz1jOfNAqTg85iguOWAvsrCJoPiezlLFA=",
        version = "v8.1.1+incompatible",
    )
    go_repository(
        name = "com_github_twitchyliquid64_golang_asm",
        importpath = "github.com/twitchyliquid64/golang-asm",
        sum = "h1:SU5vSMR7hnwNxj24w34ZyCi/FmDZTkS4MhqMhdFk5YI=",
        version = "v0.15.1",
    )
    go_repository(
        name = "com_github_u_root_uio",
        importpath = "github.com/u-root/uio",
        sum = "h1:tHNk7XK9GkmKUR6Gh8gVBKXc2MVSZ4G/NnWLtzw4gNA=",
        version = "v0.0.0-20230220225925-ffce2a382923",
    )
    go_repository(
        name = "com_github_ugorji_go",
        importpath = "github.com/ugorji/go",
        sum = "h1:/68gy2h+1mWMrwZFeD1kQialdSzAb432dtpeJ42ovdo=",
        version = "v1.1.7",
    )
    go_repository(
        name = "com_github_ugorji_go_codec",
        importpath = "github.com/ugorji/go/codec",
        sum = "h1:BMaWp1Bb6fHwEtbplGBGJ498wD+LKlNSl25MjdZY4dU=",
        version = "v1.2.11",
    )
    go_repository(
        name = "com_github_ulikunitz_xz",
        importpath = "github.com/ulikunitz/xz",
        sum = "h1:YvTNdFzX6+W5m9msiYg/zpkSURPPtOlzbqYjrFn7Yt4=",
        version = "v0.5.7",
    )
    go_repository(
        name = "com_github_urfave_cli",
        importpath = "github.com/urfave/cli",
        sum = "h1:igJgVw1JdKH+trcLWLeLwZjU9fEfPesQ+9/e4MQ44S8=",
        version = "v1.22.12",
    )
    go_repository(
        name = "com_github_urfave_negroni",
        importpath = "github.com/urfave/negroni",
        sum = "h1:kIimOitoypq34K7TG7DUaJ9kq/N4Ofuwi1sjz0KipXc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_bytebufferpool",
        importpath = "github.com/valyala/bytebufferpool",
        sum = "h1:GqA5TC/0021Y/b9FG4Oi9Mr3q7XYx6KllzawFIhcdPw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_valyala_fasthttp",
        importpath = "github.com/valyala/fasthttp",
        sum = "h1:zPkkzpIn8tdHZUrVa6PzYd0i5verqiPSkgTd3bSUcpA=",
        version = "v1.45.0",
    )
    go_repository(
        name = "com_github_valyala_fasttemplate",
        importpath = "github.com/valyala/fasttemplate",
        sum = "h1:TVEnxayobAdVkhQfrfes2IzOB6o+z4roRkPF52WA1u4=",
        version = "v1.2.1",
    )
    go_repository(
        name = "com_github_valyala_tcplisten",
        importpath = "github.com/valyala/tcplisten",
        sum = "h1:rBHj/Xf+E1tRGZyWIWwJDiRY0zc1Js+CV5DqwacVSA8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_vbatts_tar_split",
        importpath = "github.com/vbatts/tar-split",
        sum = "h1:Via6XqJr0hceW4wff3QRzD5gAk/tatMw/4ZA7cTlIME=",
        version = "v0.11.2",
    )
    go_repository(
        name = "com_github_vektah_gqlparser_v2",
        importpath = "github.com/vektah/gqlparser/v2",
        sum = "h1:C02NsyEsL4TXJB7ndonqTfuQOL4XPIu0aAWugdmTgmc=",
        version = "v2.4.5",
    )
    go_repository(
        name = "com_github_veraison_go_cose",
        importpath = "github.com/veraison/go-cose",
        sum = "h1:4qA7dbFJGvt7gcqv5MCIyCQvN+NpHFPkW7do3EeDLb8=",
        version = "v1.0.0-rc.1",
    )
    go_repository(
        name = "com_github_vishvananda_netlink",
        importpath = "github.com/vishvananda/netlink",
        patch_args = ["-p1"],
        patches = [
            # Going upstream as https://github.com/vishvananda/netlink/pull/850
            "//third_party/go/patches:netlink-permhwaddr.patch",
            "//third_party/go/patches:netlink-use-rtm_setlink.patch",
        ],
        replace = "github.com/monogon-dev/netlink",
        sum = "h1:y05BDqZ6q3if6pYBHJcnQRUd92ihzBEJde/S4fpKEAM=",
        version = "v0.0.0-20230125113930-88977c3ff4b3",
    )
    go_repository(
        name = "com_github_vishvananda_netns",
        importpath = "github.com/vishvananda/netns",
        sum = "h1:Oeaw1EM2JMxD51g9uhtC0D7erkIjgmj8+JZc26m1YX8=",
        version = "v0.0.4",
    )
    go_repository(
        name = "com_github_vividcortex_gohistogram",
        importpath = "github.com/VividCortex/gohistogram",
        sum = "h1:6+hBz+qvs0JOrrNhhmR7lFxo5sINxBCGXrdtl/UvroE=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_vmihailenco_bufpool",
        importpath = "github.com/vmihailenco/bufpool",
        sum = "h1:gOq2WmBrq0i2yW5QJ16ykccQ4wH9UyEsgLm6czKAd94=",
        version = "v0.1.11",
    )
    go_repository(
        name = "com_github_vmihailenco_msgpack_v5",
        importpath = "github.com/vmihailenco/msgpack/v5",
        sum = "h1:qMKAwOV+meBw2Y8k9cVwAy7qErtYCwBzZ2ellBfvnqc=",
        version = "v5.3.4",
    )
    go_repository(
        name = "com_github_vmihailenco_tagparser",
        importpath = "github.com/vmihailenco/tagparser",
        sum = "h1:gnjoVuB/kljJ5wICEEOpx98oXMWPLj22G67Vbd1qPqc=",
        version = "v0.1.2",
    )
    go_repository(
        name = "com_github_vmihailenco_tagparser_v2",
        importpath = "github.com/vmihailenco/tagparser/v2",
        sum = "h1:y09buUbR+b5aycVFQs/g70pqKVZNBmxwAhO7/IwNM9g=",
        version = "v2.0.0",
    )
    go_repository(
        name = "com_github_willf_bitset",
        importpath = "github.com/willf/bitset",
        sum = "h1:N7Z7E9UvjW+sGsEl7k/SJrvY2reP1A07MrGuCjIOjRE=",
        version = "v1.1.11",
    )
    go_repository(
        name = "com_github_x448_float16",
        importpath = "github.com/x448/float16",
        sum = "h1:qLwI1I70+NjRFUR3zs1JPUCgaCXSh3SW62uAKT1mSBM=",
        version = "v0.8.4",
    )
    go_repository(
        name = "com_github_xanzy_go_gitlab",
        importpath = "github.com/xanzy/go-gitlab",
        sum = "h1:+nHztQuCXGSMluKe5Q9IRaPdz6tO8O0gMkQ0vqGpiBk=",
        version = "v0.31.0",
    )
    go_repository(
        name = "com_github_xanzy_ssh_agent",
        importpath = "github.com/xanzy/ssh-agent",
        sum = "h1:TCbipTQL2JiiCprBWx9frJ2eJlCYT00NmctrHxVAr70=",
        version = "v0.2.1",
    )
    go_repository(
        name = "com_github_xdg_go_pbkdf2",
        importpath = "github.com/xdg-go/pbkdf2",
        sum = "h1:Su7DPu48wXMwC3bs7MCNG+z4FhcyEuz5dlvchbq0B0c=",
        version = "v1.0.0",
    )
    go_repository(
        name = "com_github_xdg_go_scram",
        importpath = "github.com/xdg-go/scram",
        sum = "h1:akYIkZ28e6A96dkWNJQu3nmCzH3YfwMPQExUYDaRv7w=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_xdg_go_stringprep",
        importpath = "github.com/xdg-go/stringprep",
        sum = "h1:6iq84/ryjjeRmMJwxutI51F2GIPlP5BfTvXHeYjyhBc=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonpointer",
        importpath = "github.com/xeipuuv/gojsonpointer",
        sum = "h1:zGWFAtiMcyryUHoUjUJX0/lt1H2+i2Ka2n+D3DImSNo=",
        version = "v0.0.0-20190905194746-02993c407bfb",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonreference",
        importpath = "github.com/xeipuuv/gojsonreference",
        sum = "h1:EzJWgHovont7NscjpAxXsDA8S8BMYve8Y5+7cuRE7R0=",
        version = "v0.0.0-20180127040603-bd5ef7bd5415",
    )
    go_repository(
        name = "com_github_xeipuuv_gojsonschema",
        importpath = "github.com/xeipuuv/gojsonschema",
        sum = "h1:LhYJRs+L4fBtjZUfuSZIKGeVu0QRy8e5Xi7D17UxZ74=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration",
        importpath = "github.com/xhit/go-str2duration",
        sum = "h1:BcV5u025cITWxEQKGWr1URRzrcXtu7uk8+luz3Yuhwc=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xhit_go_str2duration_v2",
        importpath = "github.com/xhit/go-str2duration/v2",
        sum = "h1:lxklc02Drh6ynqX+DdPyp5pCKLUQpRT8bp8Ydu2Bstc=",
        version = "v2.1.0",
    )
    go_repository(
        name = "com_github_xi2_xz",
        importpath = "github.com/xi2/xz",
        sum = "h1:nIPpBwaJSVYIxUFsDv3M8ofmx9yWTog9BfvIu0q41lo=",
        version = "v0.0.0-20171230120015-48954b6210f8",
    )
    go_repository(
        name = "com_github_xiang90_probing",
        importpath = "github.com/xiang90/probing",
        sum = "h1:eY9dn8+vbi4tKz5Qo6v2eYzo7kUS51QINcR5jNpbZS8=",
        version = "v0.0.0-20190116061207-43a291ad63a2",
    )
    go_repository(
        name = "com_github_xlab_treeprint",
        importpath = "github.com/xlab/treeprint",
        sum = "h1:HzHnuAF1plUN2zGlAFHbSQP2qJ0ZAD3XF5XD7OesXRQ=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_github_xordataexchange_crypt",
        importpath = "github.com/xordataexchange/crypt",
        sum = "h1:ESFSdwYZvkeru3RtdrYueztKhOBCSAAzS4Gf+k0tEow=",
        version = "v0.0.3-0.20170626215501-b2862e3d0a77",
    )
    go_repository(
        name = "com_github_yalue_native_endian",
        importpath = "github.com/yalue/native_endian",
        sum = "h1:e4SxBbaCoOOO4E3axd7FSriUhzc1bIzqZGG5jl6Evbg=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_yashtewari_glob_intersection",
        importpath = "github.com/yashtewari/glob-intersection",
        sum = "h1:6gJvMYQlTDOL3dMsPF6J0+26vwX9MB8/1q3uAdhmTrg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "com_github_youmark_pkcs8",
        importpath = "github.com/youmark/pkcs8",
        sum = "h1:splanxYIlg+5LfHAM6xpdFEAYOk8iySO56hMFq6uLyA=",
        version = "v0.0.0-20181117223130-1be2e3e5546d",
    )
    go_repository(
        name = "com_github_yuin_goldmark",
        importpath = "github.com/yuin/goldmark",
        sum = "h1:fVcFKWvrslecOb/tg+Cc05dkeYx540o0FuFt3nUVDoE=",
        version = "v1.4.13",
    )
    go_repository(
        name = "com_github_yvasiyarov_go_metrics",
        importpath = "github.com/yvasiyarov/go-metrics",
        sum = "h1:+lm10QQTNSBd8DVTNGHx7o/IKu9HYDvLMffDhbyLccI=",
        version = "v0.0.0-20140926110328-57bccd1ccd43",
    )
    go_repository(
        name = "com_github_yvasiyarov_gorelic",
        importpath = "github.com/yvasiyarov/gorelic",
        sum = "h1:hlE8//ciYMztlGpl/VA+Zm1AcTPHYkHJPbHqE6WJUXE=",
        version = "v0.0.0-20141212073537-a9bba5b9ab50",
    )
    go_repository(
        name = "com_github_yvasiyarov_newrelic_platform_go",
        importpath = "github.com/yvasiyarov/newrelic_platform_go",
        sum = "h1:ERexzlUfuTvpE74urLSbIQW0Z/6hF9t8U4NsJLaioAY=",
        version = "v0.0.0-20140908184405-b21fdbd4370f",
    )
    go_repository(
        name = "com_github_zeebo_assert",
        importpath = "github.com/zeebo/assert",
        sum = "h1:g7C04CbJuIDKNPFHmsk4hwZDO5O+kntRxzaUoNXj+IQ=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_zeebo_xxh3",
        importpath = "github.com/zeebo/xxh3",
        sum = "h1:xZmwmqxHZA8AI603jOQ0tMqmBr9lPeFwGg6d+xy9DC0=",
        version = "v1.0.2",
    )
    go_repository(
        name = "com_github_zenazn_goji",
        importpath = "github.com/zenazn/goji",
        sum = "h1:4lbD8Mx2h7IvloP7r2C0D6ltZP6Ufip8Hn0wmSK5LR8=",
        version = "v1.0.1",
    )
    go_repository(
        name = "com_gitlab_nyarla_go_crypt",
        importpath = "gitlab.com/nyarla/go-crypt",
        sum = "h1:7gd+rd8P3bqcn/96gOZa3F5dpJr/vEiDQYlNb/y2uNs=",
        version = "v0.0.0-20160106005555-d9a5dc2b789b",
    )
    go_repository(
        name = "com_google_cloud_go",
        importpath = "cloud.google.com/go",
        sum = "h1:tpFCD7hpHFlQ8yPwT3x+QeXqc2T6+n6T+hmABHfDUSM=",
        version = "v0.112.0",
    )
    go_repository(
        name = "com_google_cloud_go_accessapproval",
        importpath = "cloud.google.com/go/accessapproval",
        sum = "h1:uzmAMSgYcnlHa9X9YSQZ4Q1wlfl4NNkZyQgho1Z6p04=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_accesscontextmanager",
        importpath = "cloud.google.com/go/accesscontextmanager",
        sum = "h1:2GLNaNu9KRJhJBFTIVRoPwk6xE5mUDgD47abBq4Zp/I=",
        version = "v1.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_aiplatform",
        importpath = "cloud.google.com/go/aiplatform",
        sum = "h1:0cSrii1ZeLr16MbBoocyy5KVnrSdiQ3KN/vtrTe7RqE=",
        version = "v1.60.0",
    )
    go_repository(
        name = "com_google_cloud_go_analytics",
        importpath = "cloud.google.com/go/analytics",
        sum = "h1:Q+y94XH84jM8SK8O7qiY/PJRexb6n7dRbQ6PiUa4YGM=",
        version = "v0.23.0",
    )
    go_repository(
        name = "com_google_cloud_go_apigateway",
        importpath = "cloud.google.com/go/apigateway",
        sum = "h1:sPXnpk+6TneKIrjCjcpX5YGsAKy3PTdpIchoj8/74OE=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeconnect",
        importpath = "cloud.google.com/go/apigeeconnect",
        sum = "h1:CrfIKv9Go3fh/QfQgisU3MeP90Ww7l/sVGmr3TpECo8=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_apigeeregistry",
        importpath = "cloud.google.com/go/apigeeregistry",
        sum = "h1:C+QU2K+DzDjk4g074ouwHQGkoff1h5OMQp6sblCVreQ=",
        version = "v0.8.3",
    )
    go_repository(
        name = "com_google_cloud_go_apikeys",
        importpath = "cloud.google.com/go/apikeys",
        sum = "h1:B9CdHFZTFjVti89tmyXXrO+7vSNo2jvZuHG8zD5trdQ=",
        version = "v0.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_appengine",
        importpath = "cloud.google.com/go/appengine",
        sum = "h1:l2SviT44zWQiOv8bPoMBzW0vOcMO22iO0s+nVtVhdts=",
        version = "v1.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_area120",
        importpath = "cloud.google.com/go/area120",
        sum = "h1:vTs08KPLN/iMzTbxpu5ciL06KcsrVPMjz4IwcQyZ4uY=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_artifactregistry",
        importpath = "cloud.google.com/go/artifactregistry",
        sum = "h1:W9sVlyb1VRcUf83w7aM3yMsnp4HS4PoyGqYQNG0O5lI=",
        version = "v1.14.7",
    )
    go_repository(
        name = "com_google_cloud_go_asset",
        importpath = "cloud.google.com/go/asset",
        sum = "h1:xgFnBP3luSbUcC9RWJvb3Zkt+y/wW6PKwPHr3ssnIP8=",
        version = "v1.17.2",
    )
    go_repository(
        name = "com_google_cloud_go_assuredworkloads",
        importpath = "cloud.google.com/go/assuredworkloads",
        sum = "h1:gCrN3IyvqY3cP0wh2h43d99CgH3G+WYs9CeuFVKChR8=",
        version = "v1.11.5",
    )
    go_repository(
        name = "com_google_cloud_go_automl",
        importpath = "cloud.google.com/go/automl",
        sum = "h1:ijiJy9sYWh75WrqImXsfWc1e3HR3iO+ef9fvW03Ig/4=",
        version = "v1.13.5",
    )
    go_repository(
        name = "com_google_cloud_go_baremetalsolution",
        importpath = "cloud.google.com/go/baremetalsolution",
        sum = "h1:LFydisRmS7hQk9P/YhekwuZGqb45TW4QavcrMToWo5A=",
        version = "v1.2.4",
    )
    go_repository(
        name = "com_google_cloud_go_batch",
        importpath = "cloud.google.com/go/batch",
        sum = "h1:2HK4JerwVaIcCh/lJiHwh6+uswPthiMMWhiSWLELayk=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_beyondcorp",
        importpath = "cloud.google.com/go/beyondcorp",
        sum = "h1:qs0J0O9Ol2h1yA0AU+r7l3hOCPzs2MjE1d6d/kaHIKo=",
        version = "v1.0.4",
    )
    go_repository(
        name = "com_google_cloud_go_bigquery",
        importpath = "cloud.google.com/go/bigquery",
        sum = "h1:CpT+/njKuKT3CEmswm6IbhNu9u35zt5dO4yPDLW+nG4=",
        version = "v1.59.1",
    )
    go_repository(
        name = "com_google_cloud_go_billing",
        importpath = "cloud.google.com/go/billing",
        sum = "h1:oWUEQvuC4JvtnqLZ35zgzdbuHt4Itbftvzbe6aEyFdE=",
        version = "v1.18.2",
    )
    go_repository(
        name = "com_google_cloud_go_binaryauthorization",
        importpath = "cloud.google.com/go/binaryauthorization",
        sum = "h1:1jcyh2uIUwSZkJ/JmL8kd5SUkL/Krbv8zmYLEbAz6kY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_certificatemanager",
        importpath = "cloud.google.com/go/certificatemanager",
        sum = "h1:UMBr/twXvH3jcT5J5/YjRxf2tvwTYIfrpemTebe0txc=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_channel",
        importpath = "cloud.google.com/go/channel",
        sum = "h1:/omiBnyFjm4S1ETHoOmJbL7LH7Ljcei4rYG6Sj3hc80=",
        version = "v1.17.5",
    )
    go_repository(
        name = "com_google_cloud_go_cloudbuild",
        importpath = "cloud.google.com/go/cloudbuild",
        sum = "h1:ZB6oOmJo+MTov9n629fiCrO9YZPOg25FZvQ7gIHu5ng=",
        version = "v1.15.1",
    )
    go_repository(
        name = "com_google_cloud_go_clouddms",
        importpath = "cloud.google.com/go/clouddms",
        sum = "h1:Sr0Zo5EAcPQiCBgHWICg3VGkcdS/LLP1d9SR7qQBM/s=",
        version = "v1.7.4",
    )
    go_repository(
        name = "com_google_cloud_go_cloudtasks",
        importpath = "cloud.google.com/go/cloudtasks",
        sum = "h1:EUt1hIZ9bLv8Iz9yWaCrqgMnIU+Tdh0yXM1MMVGhjfE=",
        version = "v1.12.6",
    )
    go_repository(
        name = "com_google_cloud_go_compute",
        importpath = "cloud.google.com/go/compute",
        sum = "h1:phWcR2eWzRJaL/kOiJwfFsPs4BaKq1j6vnpZrc1YlVg=",
        version = "v1.24.0",
    )
    go_repository(
        name = "com_google_cloud_go_compute_metadata",
        importpath = "cloud.google.com/go/compute/metadata",
        sum = "h1:mg4jlk7mCAj6xXp9UJ4fjI9VUI5rubuGBW5aJ7UnBMY=",
        version = "v0.2.3",
    )
    go_repository(
        name = "com_google_cloud_go_contactcenterinsights",
        importpath = "cloud.google.com/go/contactcenterinsights",
        sum = "h1:6Vs/YnDG5STGjlWMEjN/xtmft7MrOTOnOZYUZtGTx0w=",
        version = "v1.13.0",
    )
    go_repository(
        name = "com_google_cloud_go_container",
        importpath = "cloud.google.com/go/container",
        sum = "h1:MAaNH7VRNPWEhvqOypq2j+7ONJKrKzon4v9nS3nLZe0=",
        version = "v1.31.0",
    )
    go_repository(
        name = "com_google_cloud_go_containeranalysis",
        importpath = "cloud.google.com/go/containeranalysis",
        sum = "h1:doJ0M1ljS4hS0D2UbHywlHGwB7sQLNrt9vFk9Zyi7vY=",
        version = "v0.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_datacatalog",
        importpath = "cloud.google.com/go/datacatalog",
        sum = "h1:A0vKYCQdxQuV4Pi0LL9p39Vwvg4jH5yYveMv50gU5Tw=",
        version = "v1.19.3",
    )
    go_repository(
        name = "com_google_cloud_go_dataflow",
        importpath = "cloud.google.com/go/dataflow",
        sum = "h1:RYHtcPhmE664+F0Je46p+NvFbG8z//KCXp+uEqB4jZU=",
        version = "v0.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_dataform",
        importpath = "cloud.google.com/go/dataform",
        sum = "h1:5e4eqGrd0iDTCg4Q+VlAao5j2naKAA7xRurNtwmUknU=",
        version = "v0.9.2",
    )
    go_repository(
        name = "com_google_cloud_go_datafusion",
        importpath = "cloud.google.com/go/datafusion",
        sum = "h1:HQ/BUOP8OIGJxuztpYvNvlb+/U+/Bfs9SO8tQbh61fk=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_datalabeling",
        importpath = "cloud.google.com/go/datalabeling",
        sum = "h1:GpIFRdm0qIZNsxqURFJwHt0ZBJZ0nF/mUVEigR7PH/8=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_dataplex",
        importpath = "cloud.google.com/go/dataplex",
        sum = "h1:fxIfdU8fxzR3clhOoNI7XFppvAmndxDu1AMH+qX9WKQ=",
        version = "v1.14.2",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc",
        importpath = "cloud.google.com/go/dataproc",
        sum = "h1:W47qHL3W4BPkAIbk4SWmIERwsWBaNnWm0P2sdx3YgGU=",
        version = "v1.12.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataproc_v2",
        importpath = "cloud.google.com/go/dataproc/v2",
        sum = "h1:/u81Fd+BvCLp+xjctI1DiWVJn6cn9/s3Akc8xPH02yk=",
        version = "v2.4.0",
    )
    go_repository(
        name = "com_google_cloud_go_dataqna",
        importpath = "cloud.google.com/go/dataqna",
        sum = "h1:9ybXs3nr9BzxSGC04SsvtuXaHY0qmJSLIpIAbZo9GqQ=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_datastore",
        importpath = "cloud.google.com/go/datastore",
        sum = "h1:0P9WcsQeTWjuD1H14JIY7XQscIPQ4Laje8ti96IC5vg=",
        version = "v1.15.0",
    )
    go_repository(
        name = "com_google_cloud_go_datastream",
        importpath = "cloud.google.com/go/datastream",
        sum = "h1:o1QDKMo/hk0FN7vhoUQURREuA0rgKmnYapB+1M+7Qz4=",
        version = "v1.10.4",
    )
    go_repository(
        name = "com_google_cloud_go_deploy",
        importpath = "cloud.google.com/go/deploy",
        sum = "h1:m27Ojwj03gvpJqCbodLYiVmE9x4/LrHGGMjzc0LBfM4=",
        version = "v1.17.1",
    )
    go_repository(
        name = "com_google_cloud_go_dialogflow",
        importpath = "cloud.google.com/go/dialogflow",
        sum = "h1:KqG0oxGE71qo0lRVyAoeBozefCvsMfcDzDjoLYSY0F4=",
        version = "v1.49.0",
    )
    go_repository(
        name = "com_google_cloud_go_dlp",
        importpath = "cloud.google.com/go/dlp",
        sum = "h1:lTipOuJaSjlYnnotPMbEhKURLC6GzCMDDzVbJAEbmYM=",
        version = "v1.11.2",
    )
    go_repository(
        name = "com_google_cloud_go_documentai",
        importpath = "cloud.google.com/go/documentai",
        sum = "h1:lI62GMEEPO6vXJI9hj+G9WjOvnR0hEjvjokrnex4cxA=",
        version = "v1.25.0",
    )
    go_repository(
        name = "com_google_cloud_go_domains",
        importpath = "cloud.google.com/go/domains",
        sum = "h1:Mml/R6s3vQQvFPpi/9oX3O5dRirgjyJ8cksK8N19Y7g=",
        version = "v0.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_edgecontainer",
        importpath = "cloud.google.com/go/edgecontainer",
        sum = "h1:tBY32km78ScpK2aOP84JoW/+wtpx5WluyPUSEE3270U=",
        version = "v1.1.5",
    )
    go_repository(
        name = "com_google_cloud_go_errorreporting",
        importpath = "cloud.google.com/go/errorreporting",
        sum = "h1:kj1XEWMu8P0qlLhm3FwcaFsUvXChV/OraZwA70trRR0=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_essentialcontacts",
        importpath = "cloud.google.com/go/essentialcontacts",
        sum = "h1:13eHn5qBnsawxI7mIrv4jRIEmQ1xg0Ztqw5ZGqtUNfA=",
        version = "v1.6.6",
    )
    go_repository(
        name = "com_google_cloud_go_eventarc",
        importpath = "cloud.google.com/go/eventarc",
        sum = "h1:ORkd6/UV5FIdA8KZQDLNZYKS7BBOrj0p01DXPmT4tE4=",
        version = "v1.13.4",
    )
    go_repository(
        name = "com_google_cloud_go_filestore",
        importpath = "cloud.google.com/go/filestore",
        sum = "h1:X5G4y/vrUo1B8Nsz93qSWTMAcM8LXbGUldq33OdcdCw=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_firestore",
        importpath = "cloud.google.com/go/firestore",
        sum = "h1:8aLcKnMPoldYU3YHgu4t2exrKhLQkqaXAGqT0ljrFVw=",
        version = "v1.14.0",
    )
    go_repository(
        name = "com_google_cloud_go_functions",
        importpath = "cloud.google.com/go/functions",
        sum = "h1:IWVylmK5F6hJ3R5zaRW7jI5PrWhCvtBVU4axQLmXSo4=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_gaming",
        importpath = "cloud.google.com/go/gaming",
        sum = "h1:5qZmZEWzMf8GEFgm9NeC3bjFRpt7x4S6U7oLbxaf7N8=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_gkebackup",
        importpath = "cloud.google.com/go/gkebackup",
        sum = "h1:iuE8KNtTsPOc79qeWoNS8zOWoXPD9SAdOmwgxtlCmh8=",
        version = "v1.3.5",
    )
    go_repository(
        name = "com_google_cloud_go_gkeconnect",
        importpath = "cloud.google.com/go/gkeconnect",
        sum = "h1:17d+ZSSXKqG/RwZCq3oFMIWLPI8Zw3b8+a9/BEVlwH0=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_gkehub",
        importpath = "cloud.google.com/go/gkehub",
        sum = "h1:RboLNFzf9wEMSo7DrKVBlf+YhK/A/jrLN454L5Tz99Q=",
        version = "v0.14.5",
    )
    go_repository(
        name = "com_google_cloud_go_gkemulticloud",
        importpath = "cloud.google.com/go/gkemulticloud",
        sum = "h1:rsSZAGLhyjyE/bE2ToT5fqo1qSW7S+Ubsc9jFOcbhSI=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_google_cloud_go_grafeas",
        importpath = "cloud.google.com/go/grafeas",
        sum = "h1:oyTL/KjiUeBs9eYLw/40cpSZglUC+0F7X4iu/8t7NWs=",
        version = "v0.3.0",
    )
    go_repository(
        name = "com_google_cloud_go_gsuiteaddons",
        importpath = "cloud.google.com/go/gsuiteaddons",
        sum = "h1:CZEbaBwmbYdhFw21Fwbo+C35HMe36fTE0FBSR4KSfWg=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_iam",
        importpath = "cloud.google.com/go/iam",
        sum = "h1:bEa06k05IO4f4uJonbB5iAgKTPpABy1ayxaIZV/GHVc=",
        version = "v1.1.6",
    )
    go_repository(
        name = "com_google_cloud_go_iap",
        importpath = "cloud.google.com/go/iap",
        sum = "h1:94zirc2r4t6KzhAMW0R6Dme005eTP6yf7g6vN4IhRrA=",
        version = "v1.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_ids",
        importpath = "cloud.google.com/go/ids",
        sum = "h1:xd4U7pgl3GHV+MABnv1BF4/Vy/zBF7CYC8XngkOLzag=",
        version = "v1.4.5",
    )
    go_repository(
        name = "com_google_cloud_go_iot",
        importpath = "cloud.google.com/go/iot",
        sum = "h1:munTeBlbqI33iuTYgXy7S8lW2TCgi5l1hA4roSIY+EE=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_kms",
        importpath = "cloud.google.com/go/kms",
        sum = "h1:7caV9K3yIxvlQPAcaFffhlT7d1qpxjB1wHBtjWa13SM=",
        version = "v1.15.7",
    )
    go_repository(
        name = "com_google_cloud_go_language",
        importpath = "cloud.google.com/go/language",
        sum = "h1:iaJZg6K4j/2PvZZVcjeO/btcWWIllVRBhuTFjGO4LXs=",
        version = "v1.12.3",
    )
    go_repository(
        name = "com_google_cloud_go_lifesciences",
        importpath = "cloud.google.com/go/lifesciences",
        sum = "h1:gXvN70m2p+4zgJFzaz6gMKaxTuF9WJ0USYoMLWAOm8g=",
        version = "v0.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_logging",
        importpath = "cloud.google.com/go/logging",
        sum = "h1:iEIOXFO9EmSiTjDmfpbRjOxECO7R8C7b8IXUGOj7xZw=",
        version = "v1.9.0",
    )
    go_repository(
        name = "com_google_cloud_go_longrunning",
        importpath = "cloud.google.com/go/longrunning",
        sum = "h1:GOE6pZFdSrTb4KAiKnXsJBtlE6mEyaW44oKyMILWnOg=",
        version = "v0.5.5",
    )
    go_repository(
        name = "com_google_cloud_go_managedidentities",
        importpath = "cloud.google.com/go/managedidentities",
        sum = "h1:+bpih1piZVLxla/XBqeSUzJBp8gv9plGHIMAI7DLpDM=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_maps",
        importpath = "cloud.google.com/go/maps",
        sum = "h1:EVCZAiDvog9So46460BGbCasPhi613exoaQbpilMVlk=",
        version = "v1.6.4",
    )
    go_repository(
        name = "com_google_cloud_go_mediatranslation",
        importpath = "cloud.google.com/go/mediatranslation",
        sum = "h1:c76KdIXljQHSCb/Cy47S8H4s05A4zbK3pAFGzwcczZo=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_memcache",
        importpath = "cloud.google.com/go/memcache",
        sum = "h1:yeDv5qxRedFosvpMSEswrqUsJM5OdWvssPHFliNFTc4=",
        version = "v1.10.5",
    )
    go_repository(
        name = "com_google_cloud_go_metastore",
        importpath = "cloud.google.com/go/metastore",
        sum = "h1:dR7vqWXlK6IYR8Wbu9mdFfwlVjodIBhd1JRrpZftTEg=",
        version = "v1.13.4",
    )
    go_repository(
        name = "com_google_cloud_go_monitoring",
        importpath = "cloud.google.com/go/monitoring",
        sum = "h1:NfkDLQDG2UR3WYZVQE8kwSbUIEyIqJUPl+aOQdFH1T4=",
        version = "v1.18.0",
    )
    go_repository(
        name = "com_google_cloud_go_networkconnectivity",
        importpath = "cloud.google.com/go/networkconnectivity",
        sum = "h1:GBfXFhLyPspnaBE3nI/BRjdhW8vcbpT9QjE/4kDCDdc=",
        version = "v1.14.4",
    )
    go_repository(
        name = "com_google_cloud_go_networkmanagement",
        importpath = "cloud.google.com/go/networkmanagement",
        sum = "h1:aLV5GcosBNmd6M8+a0ekB0XlLRexv4fvnJJrYnqeBcg=",
        version = "v1.9.4",
    )
    go_repository(
        name = "com_google_cloud_go_networksecurity",
        importpath = "cloud.google.com/go/networksecurity",
        sum = "h1:+caSxBTj0E8OYVh/5wElFdjEMO1S/rZtE1152Cepchc=",
        version = "v0.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_notebooks",
        importpath = "cloud.google.com/go/notebooks",
        sum = "h1:FH48boYmrWVQ6k0Mx/WrnNafXncT5iSYxA8CNyWTgy0=",
        version = "v1.11.3",
    )
    go_repository(
        name = "com_google_cloud_go_optimization",
        importpath = "cloud.google.com/go/optimization",
        sum = "h1:63NZaWyN+5rZEKHPX4ACpw3BjgyeuY8+rCehiCMaGPY=",
        version = "v1.6.3",
    )
    go_repository(
        name = "com_google_cloud_go_orchestration",
        importpath = "cloud.google.com/go/orchestration",
        sum = "h1:YHgWMlrPttIVGItgGfuvO2KM7x+y9ivN/Yk92pMm1a4=",
        version = "v1.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_orgpolicy",
        importpath = "cloud.google.com/go/orgpolicy",
        sum = "h1:2JbXigqBJVp8Dx5dONUttFqewu4fP0p3pgOdIZAhpYU=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_osconfig",
        importpath = "cloud.google.com/go/osconfig",
        sum = "h1:Mo5jGAxOMKH/PmDY7fgY19yFcVbvwREb5D5zMPQjFfo=",
        version = "v1.12.5",
    )
    go_repository(
        name = "com_google_cloud_go_oslogin",
        importpath = "cloud.google.com/go/oslogin",
        sum = "h1:1K4nOT5VEZNt7XkhaTXupBYos5HjzvJMfhvyD2wWdFs=",
        version = "v1.13.1",
    )
    go_repository(
        name = "com_google_cloud_go_phishingprotection",
        importpath = "cloud.google.com/go/phishingprotection",
        sum = "h1:DH3WFLzEoJdW/6xgsmoDqOwT1xddFi7gKu0QGZQhpGU=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_policytroubleshooter",
        importpath = "cloud.google.com/go/policytroubleshooter",
        sum = "h1:c0WOzC6hz964QWNBkyKfna8A2jOIx1zzZa43Gx/P09o=",
        version = "v1.10.3",
    )
    go_repository(
        name = "com_google_cloud_go_privatecatalog",
        importpath = "cloud.google.com/go/privatecatalog",
        sum = "h1:UZ0assTnATXSggoxUIh61RjTQ4P9zCMk/kEMbn0nMYA=",
        version = "v0.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_pubsub",
        importpath = "cloud.google.com/go/pubsub",
        sum = "h1:dfEPuGCHGbWUhaMCTHUFjfroILEkx55iUmKBZTP5f+Y=",
        version = "v1.36.1",
    )
    go_repository(
        name = "com_google_cloud_go_pubsublite",
        importpath = "cloud.google.com/go/pubsublite",
        sum = "h1:pX+idpWMIH30/K7c0epN6V703xpIcMXWRjKJsz0tYGY=",
        version = "v1.8.1",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise",
        importpath = "cloud.google.com/go/recaptchaenterprise",
        sum = "h1:u6EznTGzIdsyOsvm+Xkw0aSuKFXQlyjGE9a4exk6iNQ=",
        version = "v1.3.1",
    )
    go_repository(
        name = "com_google_cloud_go_recaptchaenterprise_v2",
        importpath = "cloud.google.com/go/recaptchaenterprise/v2",
        sum = "h1:U3Wfq12X9cVMuTpsWDSURnXF0Z9hSPTHj+xsnXDRLsw=",
        version = "v2.9.2",
    )
    go_repository(
        name = "com_google_cloud_go_recommendationengine",
        importpath = "cloud.google.com/go/recommendationengine",
        sum = "h1:ineqLswaCSBY0csYv5/wuXJMBlxATK6Xc5jJkpiTEdM=",
        version = "v0.8.5",
    )
    go_repository(
        name = "com_google_cloud_go_recommender",
        importpath = "cloud.google.com/go/recommender",
        sum = "h1:LVLYS3r3u0MSCxQSDUtLSkporEGi9OAE6hGvayrZNPs=",
        version = "v1.12.1",
    )
    go_repository(
        name = "com_google_cloud_go_redis",
        importpath = "cloud.google.com/go/redis",
        sum = "h1:QF0maEdVv0Fj/2roU8sX3NpiDBzP9ICYTO+5F32gQNo=",
        version = "v1.14.2",
    )
    go_repository(
        name = "com_google_cloud_go_resourcemanager",
        importpath = "cloud.google.com/go/resourcemanager",
        sum = "h1:AZWr1vWVDKGwfLsVhcN+vcwOz3xqqYxtmMa0aABCMms=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_resourcesettings",
        importpath = "cloud.google.com/go/resourcesettings",
        sum = "h1:BTr5MVykJwClASci/7Og4Qfx70aQ4n3epsNLj94ZYgw=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_retail",
        importpath = "cloud.google.com/go/retail",
        sum = "h1:Fn1GuAua1c6crCGqfJ1qMxG1Xh10Tg/x5EUODEHMqkw=",
        version = "v1.16.0",
    )
    go_repository(
        name = "com_google_cloud_go_run",
        importpath = "cloud.google.com/go/run",
        sum = "h1:m9WDA7DzTpczhZggwYlZcBWgCRb+kgSIisWn1sbw2rQ=",
        version = "v1.3.4",
    )
    go_repository(
        name = "com_google_cloud_go_scheduler",
        importpath = "cloud.google.com/go/scheduler",
        sum = "h1:5U8iXLoQ03qOB+ZXlAecU7fiE33+u3QiM9nh4cd0eTE=",
        version = "v1.10.6",
    )
    go_repository(
        name = "com_google_cloud_go_secretmanager",
        importpath = "cloud.google.com/go/secretmanager",
        sum = "h1:82fpF5vBBvu9XW4qj0FU2C6qVMtj1RM/XHwKXUEAfYY=",
        version = "v1.11.5",
    )
    go_repository(
        name = "com_google_cloud_go_security",
        importpath = "cloud.google.com/go/security",
        sum = "h1:wTKJQ10j8EYgvE8Y+KhovxDRVDk2iv/OsxZ6GrLP3kE=",
        version = "v1.15.5",
    )
    go_repository(
        name = "com_google_cloud_go_securitycenter",
        importpath = "cloud.google.com/go/securitycenter",
        sum = "h1:/5jjkZ+uGe8hZ7pvd7pO30VW/a+pT2MrrdgOqjyucKQ=",
        version = "v1.24.4",
    )
    go_repository(
        name = "com_google_cloud_go_servicecontrol",
        importpath = "cloud.google.com/go/servicecontrol",
        sum = "h1:d0uV7Qegtfaa7Z2ClDzr9HJmnbJW7jn0WhZ7wOX6hLE=",
        version = "v1.11.1",
    )
    go_repository(
        name = "com_google_cloud_go_servicedirectory",
        importpath = "cloud.google.com/go/servicedirectory",
        sum = "h1:da7HFI1229kyzIyuVEzHXip0cw0d+E0s8mjQby0WN+k=",
        version = "v1.11.4",
    )
    go_repository(
        name = "com_google_cloud_go_servicemanagement",
        importpath = "cloud.google.com/go/servicemanagement",
        sum = "h1:fopAQI/IAzlxnVeiKn/8WiV6zKndjFkvi+gzu+NjywY=",
        version = "v1.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_serviceusage",
        importpath = "cloud.google.com/go/serviceusage",
        sum = "h1:rXyq+0+RSIm3HFypctp7WoXxIA563rn206CfMWdqXX4=",
        version = "v1.6.0",
    )
    go_repository(
        name = "com_google_cloud_go_shell",
        importpath = "cloud.google.com/go/shell",
        sum = "h1:3Fq2hzO0ZSyaqBboJrFkwwf/qMufDtqwwA6ep8EZxEI=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_spanner",
        importpath = "cloud.google.com/go/spanner",
        sum = "h1:fJq+ZfQUDHE+cy1li0bJA8+sy2oiSGhuGqN5nqVaZdU=",
        version = "v1.57.0",
    )
    go_repository(
        name = "com_google_cloud_go_speech",
        importpath = "cloud.google.com/go/speech",
        sum = "h1:nuFc+Kj5B8de75nN4FdPyUbI2SiBoHZG6BLurXL56Q0=",
        version = "v1.21.1",
    )
    go_repository(
        name = "com_google_cloud_go_storage",
        importpath = "cloud.google.com/go/storage",
        sum = "h1:P0mOkAcaJxhCTvAkMhxMfrTKiNcub4YmmPBtlhAyTr8=",
        version = "v1.36.0",
    )
    go_repository(
        name = "com_google_cloud_go_storagetransfer",
        importpath = "cloud.google.com/go/storagetransfer",
        sum = "h1:dy4fL3wO0VABvzM05ycMUPFHxTPbJz9Em8ikAJVqSbI=",
        version = "v1.10.4",
    )
    go_repository(
        name = "com_google_cloud_go_talent",
        importpath = "cloud.google.com/go/talent",
        sum = "h1:JssV0CE3FNujuSWn7SkosOzg7qrMxVnt6txOfGcMSa4=",
        version = "v1.6.6",
    )
    go_repository(
        name = "com_google_cloud_go_texttospeech",
        importpath = "cloud.google.com/go/texttospeech",
        sum = "h1:dxY2Q5mHCbrGa3oPR2O3PCicdnvKa1JmwGQK36EFLOw=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_tpu",
        importpath = "cloud.google.com/go/tpu",
        sum = "h1:C8YyYda8WtNdBoCgFwwBzZd+S6+EScHOxM/z1h0NNp8=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_trace",
        importpath = "cloud.google.com/go/trace",
        sum = "h1:0pr4lIKJ5XZFYD9GtxXEWr0KkVeigc3wlGpZco0X1oA=",
        version = "v1.10.5",
    )
    go_repository(
        name = "com_google_cloud_go_translate",
        importpath = "cloud.google.com/go/translate",
        sum = "h1:upovZ0wRMdzZvXnu+RPam41B0mRJ+coRXFP2cYFJ7ew=",
        version = "v1.10.1",
    )
    go_repository(
        name = "com_google_cloud_go_video",
        importpath = "cloud.google.com/go/video",
        sum = "h1:TXwotxkShP1OqgKsbd+b8N5hrIHavSyLGvYnLGCZ7xc=",
        version = "v1.20.4",
    )
    go_repository(
        name = "com_google_cloud_go_videointelligence",
        importpath = "cloud.google.com/go/videointelligence",
        sum = "h1:mYaWH8uhUCXLJCN3gdXswKzRa2+lK0zN6/KsIubm6pE=",
        version = "v1.11.5",
    )
    go_repository(
        name = "com_google_cloud_go_vision",
        importpath = "cloud.google.com/go/vision",
        sum = "h1:/CsSTkbmO9HC8iQpxbK8ATms3OQaX3YQUeTMGCxlaK4=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_google_cloud_go_vision_v2",
        importpath = "cloud.google.com/go/vision/v2",
        sum = "h1:W52z1b6LdGI66MVhE70g/NFty9zCYYcjdKuycqmlhtg=",
        version = "v2.8.0",
    )
    go_repository(
        name = "com_google_cloud_go_vmmigration",
        importpath = "cloud.google.com/go/vmmigration",
        sum = "h1:5v9RT2vWyuw3pK2ox0HQpkoftO7Q7/8591dTxxQc79g=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_vmwareengine",
        importpath = "cloud.google.com/go/vmwareengine",
        sum = "h1:EGdDi9QbqThfZq3ILcDK5g+m9jTevc34AY5tACx5v7k=",
        version = "v1.1.1",
    )
    go_repository(
        name = "com_google_cloud_go_vpcaccess",
        importpath = "cloud.google.com/go/vpcaccess",
        sum = "h1:XyL6hTLtEM/eE4F1GEge8xUN9ZCkiVWn44K/YA7z1rQ=",
        version = "v1.7.5",
    )
    go_repository(
        name = "com_google_cloud_go_webrisk",
        importpath = "cloud.google.com/go/webrisk",
        sum = "h1:251MvGuC8wisNN7+jqu9DDDZAi38KiMXxOpA/EWy4dE=",
        version = "v1.9.5",
    )
    go_repository(
        name = "com_google_cloud_go_websecurityscanner",
        importpath = "cloud.google.com/go/websecurityscanner",
        sum = "h1:YqWZrZYabG88TZt7364XWRJGhxmxhony2ZUyZEYMF2k=",
        version = "v1.6.5",
    )
    go_repository(
        name = "com_google_cloud_go_workflows",
        importpath = "cloud.google.com/go/workflows",
        sum = "h1:uHNmUiatTbPQ4H1pabwfzpfEYD4BBnqDHqMm2IesOh4=",
        version = "v1.12.4",
    )
    go_repository(
        name = "com_lukechampine_uint128",
        importpath = "lukechampine.com/uint128",
        sum = "h1:mBi/5l91vocEN8otkC5bDLhi2KdCticRiwbdB0O+rjI=",
        version = "v1.2.0",
    )
    go_repository(
        name = "com_sourcegraph_sourcegraph_appdash",
        importpath = "sourcegraph.com/sourcegraph/appdash",
        sum = "h1:ucqkfpjg9WzSUubAO62csmucvxl4/JeW3F4I4909XkM=",
        version = "v0.0.0-20190731080439-ebfcffb1b5c0",
    )
    go_repository(
        name = "com_zx2c4_golang_go118_netip",
        importpath = "golang.zx2c4.com/go118/netip",
        sum = "h1:9+v0G0naRhLPOJEeJOL6NuXTtAHHwmkyZlgQJ0XcQ8I=",
        version = "v0.0.0-20211111135330-a4a02eeacf9d",
    )
    go_repository(
        name = "com_zx2c4_golang_wintun",
        importpath = "golang.zx2c4.com/wintun",
        sum = "h1:Ug9qvr1myri/zFN6xL17LSCBGFDnphBBhzmILHsM5TY=",
        version = "v0.0.0-20211104114900-415007cec224",
    )
    go_repository(
        name = "com_zx2c4_golang_wireguard",
        importpath = "golang.zx2c4.com/wireguard",
        sum = "h1:Nrf94TOjrvW8nm6N3u2xtbnMZaZudNI9b8nIJH8p8qY=",
        version = "v0.0.0-20220202223031-3b95c81cc178",
    )
    go_repository(
        name = "com_zx2c4_golang_wireguard_wgctrl",
        importpath = "golang.zx2c4.com/wireguard/wgctrl",
        sum = "h1:9c4/JVIQUc2qCJEEIiGIs3HmmnFjhPj4qHW4+Uj+u3U=",
        version = "v0.0.0-20220208144051-fde48d68ee68",
    )
    go_repository(
        name = "dev_gocloud",
        importpath = "gocloud.dev",
        sum = "h1:EDRyaRAnMGSq/QBto486gWFxMLczAfIYUmusV7XLNBM=",
        version = "v0.19.0",
    )
    go_repository(
        name = "dev_gvisor_gvisor",
        importpath = "gvisor.dev/gvisor",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:gvisor-fix-resolution.patch",
            "//third_party/go/patches:gvisor-fix-syslog.patch",
            # Gazelle doesn't understand conditional build files.
            # https://github.com/bazelbuild/bazel-gazelle/issues/1262
            "//third_party/go/patches:gvisor-fix-gazelle-debug-builds.patch",
        ],
        pre_patches = [
            "//third_party/go/patches:gvisor-containerd-compat.patch",
            # Modified patch for the go branch.
            # Upstream issue https://github.com/google/gvisor/issues/9039
            # Upstream patch https://github.com/google/gvisor/commit/62afa160fe8ad17302d43f7c6016061de57c7f4d
            "//third_party/go/patches:gvisor-fix-debug-builds.patch",
        ],
        sum = "h1:jkIUiFfucTr4prtWuxQ1YxuOdWEQc/dIrUnHKyzSQv0=",
        version = "v0.0.0-20230911190645-2e1d76499fd5",
    )
    go_repository(
        name = "ht_sr_git_sbinet_gg",
        importpath = "git.sr.ht/~sbinet/gg",
        sum = "h1:LNhjNn8DerC8f9DHLz6lS0YYul/b602DUxDgGkd/Aik=",
        version = "v0.3.1",
    )
    go_repository(
        name = "im_mellium_sasl",
        importpath = "mellium.im/sasl",
        sum = "h1:wE0LW6g7U83vhvxjC1IY8DnXM+EU095yeo8XClvCdfo=",
        version = "v0.3.1",
    )
    go_repository(
        name = "in_gopkg_airbrake_gobrake_v2",
        importpath = "gopkg.in/airbrake/gobrake.v2",
        sum = "h1:7z2uVWwn7oVeeugY1DtlPAy5H+KYgB1KeKTnqjNatLo=",
        version = "v2.0.9",
    )
    go_repository(
        name = "in_gopkg_alecthomas_kingpin_v2",
        importpath = "gopkg.in/alecthomas/kingpin.v2",
        sum = "h1:jMFz6MfLP0/4fUyZle81rXUoxOBFi19VUFKVDOQfozc=",
        version = "v2.2.6",
    )
    go_repository(
        name = "in_gopkg_check_v1",
        importpath = "gopkg.in/check.v1",
        sum = "h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=",
        version = "v1.0.0-20201130134442-10cb98267c6c",
    )
    go_repository(
        name = "in_gopkg_cheggaaa_pb_v1",
        importpath = "gopkg.in/cheggaaa/pb.v1",
        sum = "h1:n1tBJnnK2r7g9OW2btFH91V92STTUevLXYFb8gy9EMk=",
        version = "v1.0.28",
    )
    go_repository(
        name = "in_gopkg_datadog_dd_trace_go_v1",
        importpath = "gopkg.in/DataDog/dd-trace-go.v1",
        sum = "h1:LAFmtVYLnqhsFAsKv3799SYalXD9Hl3K0/pR+3eV/Qc=",
        version = "v1.54.0",
    )
    go_repository(
        name = "in_gopkg_djherbis_times_v1",
        importpath = "gopkg.in/djherbis/times.v1",
        sum = "h1:UCvDKl1L/fmBygl2Y7hubXCnY7t4Yj46ZrBFNUipFbM=",
        version = "v1.2.0",
    )
    go_repository(
        name = "in_gopkg_errgo_v2",
        importpath = "gopkg.in/errgo.v2",
        sum = "h1:0vLT13EuvQ0hNvakwLuFZ/jYrLp5F3kcWHXdRggjCE8=",
        version = "v2.1.0",
    )
    go_repository(
        name = "in_gopkg_fsnotify_v1",
        importpath = "gopkg.in/fsnotify.v1",
        sum = "h1:xOHLXZwVvI9hhs+cLKq5+I5onOuwQLhQwiu63xxlHs4=",
        version = "v1.4.7",
    )
    go_repository(
        name = "in_gopkg_gcfg_v1",
        importpath = "gopkg.in/gcfg.v1",
        sum = "h1:m8OOJ4ccYHnx2f4gQwpno8nAX5OGOh7RLaaz0pj3Ogs=",
        version = "v1.2.3",
    )
    go_repository(
        name = "in_gopkg_gemnasium_logrus_airbrake_hook_v2",
        importpath = "gopkg.in/gemnasium/logrus-airbrake-hook.v2",
        sum = "h1:OAj3g0cR6Dx/R07QgQe8wkA9RNjB2u4i700xBkIT4e0=",
        version = "v2.1.2",
    )
    go_repository(
        name = "in_gopkg_go_playground_assert_v1",
        importpath = "gopkg.in/go-playground/assert.v1",
        sum = "h1:xoYuJVE7KT85PYWrN730RguIQO0ePzVRfFMXadIrXTM=",
        version = "v1.2.1",
    )
    go_repository(
        name = "in_gopkg_go_playground_validator_v9",
        importpath = "gopkg.in/go-playground/validator.v9",
        sum = "h1:SvGtYmN60a5CVKTOzMSyfzWDeZRxRuGvRQyEAKbw1xc=",
        version = "v9.29.1",
    )
    go_repository(
        name = "in_gopkg_inconshreveable_log15_v2",
        importpath = "gopkg.in/inconshreveable/log15.v2",
        sum = "h1:RlWgLqCMMIYYEVcAR5MDsuHlVkaIPDAF+5Dehzg8L5A=",
        version = "v2.0.0-20180818164646-67afb5ed74ec",
    )
    go_repository(
        name = "in_gopkg_inf_v0",
        importpath = "gopkg.in/inf.v0",
        sum = "h1:73M5CoZyi3ZLMOyDlQh031Cx6N9NDJ2Vvfl76EDAgDc=",
        version = "v0.9.1",
    )
    go_repository(
        name = "in_gopkg_ini_v1",
        importpath = "gopkg.in/ini.v1",
        sum = "h1:AQvPpx3LzTDM0AjnIRlVFwFFGC+npRopjZxLJj6gdno=",
        version = "v1.51.0",
    )
    go_repository(
        name = "in_gopkg_jinzhu_gorm_v1",
        importpath = "gopkg.in/jinzhu/gorm.v1",
        sum = "h1:63D1Sk0C0mhCbK930D0PkD3nKT8wLxz6lLPh5V6D2hM=",
        version = "v1.9.1",
    )
    go_repository(
        name = "in_gopkg_natefinch_lumberjack_v2",
        importpath = "gopkg.in/natefinch/lumberjack.v2",
        sum = "h1:bBRl1b0OH9s/DuPhuXpNl+VtCaJXFZ5/uEFST95x9zc=",
        version = "v2.2.1",
    )
    go_repository(
        name = "in_gopkg_olivere_elastic_v3",
        importpath = "gopkg.in/olivere/elastic.v3",
        sum = "h1:u3B8p1VlHF3yNLVOlhIWFT3F1ICcHfM5V6FFJe6pPSo=",
        version = "v3.0.75",
    )
    go_repository(
        name = "in_gopkg_olivere_elastic_v5",
        importpath = "gopkg.in/olivere/elastic.v5",
        sum = "h1:acF/tRSg5geZpE3rqLglkS79CQMIMzOpWZE7hRXIkjs=",
        version = "v5.0.84",
    )
    go_repository(
        name = "in_gopkg_resty_v1",
        importpath = "gopkg.in/resty.v1",
        sum = "h1:CuXP0Pjfw9rOuY6EP+UvtNvt5DSqHpIxILZKT/quCZI=",
        version = "v1.12.0",
    )
    go_repository(
        name = "in_gopkg_square_go_jose_v2",
        importpath = "gopkg.in/square/go-jose.v2",
        sum = "h1:NGk74WTnPKBNUhNzQX7PYcTLUjoq7mzKk2OKbvwk2iI=",
        version = "v2.6.0",
    )
    go_repository(
        name = "in_gopkg_src_d_go_billy_v4",
        importpath = "gopkg.in/src-d/go-billy.v4",
        sum = "h1:0SQA1pRztfTFx2miS8sA97XvooFeNOmvUenF4o0EcVg=",
        version = "v4.3.2",
    )
    go_repository(
        name = "in_gopkg_src_d_go_git_fixtures_v3",
        importpath = "gopkg.in/src-d/go-git-fixtures.v3",
        sum = "h1:ivZFOIltbce2Mo8IjzUHAFoq/IylO9WHhNOAJK+LsJg=",
        version = "v3.5.0",
    )
    go_repository(
        name = "in_gopkg_src_d_go_git_v4",
        importpath = "gopkg.in/src-d/go-git.v4",
        sum = "h1:SRtFyV8Kxc0UP7aCHcijOMQGPxHSmMOPrzulQWolkYE=",
        version = "v4.13.1",
    )
    go_repository(
        name = "in_gopkg_tomb_v1",
        importpath = "gopkg.in/tomb.v1",
        sum = "h1:uRGJdciOHaEIrze2W8Q3AKkepLTh2hOroT7a+7czfdQ=",
        version = "v1.0.0-20141024135613-dd632973f1e7",
    )
    go_repository(
        name = "in_gopkg_warnings_v0",
        importpath = "gopkg.in/warnings.v0",
        sum = "h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=",
        version = "v0.1.2",
    )
    go_repository(
        name = "in_gopkg_yaml_v1",
        importpath = "gopkg.in/yaml.v1",
        sum = "h1:POO/ycCATvegFmVuPpQzZFJ+pGZeX22Ufu6fibxDVjU=",
        version = "v1.0.0-20140924161607-9f9df34309c0",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        sum = "h1:D8xgwECY7CYvx+Y2n4sBz93Jn9JRvxdiyyo8CTfuKaY=",
        version = "v2.4.0",
    )
    go_repository(
        name = "in_gopkg_yaml_v3",
        importpath = "gopkg.in/yaml.v3",
        sum = "h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=",
        version = "v3.0.1",
    )
    go_repository(
        name = "io_cncf_tags_container_device_interface",
        importpath = "tags.cncf.io/container-device-interface",
        sum = "h1:dThE6dtp/93ZDGhqaED2Pu374SOeUkBfuvkLuiTdwzg=",
        version = "v0.6.2",
    )
    go_repository(
        name = "io_cncf_tags_container_device_interface_specs_go",
        importpath = "tags.cncf.io/container-device-interface/specs-go",
        sum = "h1:V+tJJN6dqu8Vym6p+Ru+K5mJ49WL6Aoc5SJFSY0RLsQ=",
        version = "v0.6.0",
    )
    go_repository(
        name = "io_etcd_go_bbolt",
        importpath = "go.etcd.io/bbolt",
        sum = "h1:8x7aARPEXiXbHmtUwAIv7eV2fQFHrLLavdiJ3uzJXoI=",
        version = "v1.3.9",
    )
    go_repository(
        name = "io_etcd_go_etcd",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd",
        sum = "h1:/RwRVN9EdXAVtdHxP7Ndn/tfmM9/goiwU0QTnLBgS4w=",
        version = "v0.0.0-20200513171258-e048e166ab9c",
    )
    go_repository(
        name = "io_etcd_go_etcd_api_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/api/v3",
        sum = "h1:8WXU2/NBge6AUF1K1gOexB6e07NgsN1hXK0rSTtgSp4=",
        version = "v3.5.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_pkg_v3",
        importpath = "go.etcd.io/etcd/client/pkg/v3",
        sum = "h1:RVZSAnWWWiI5IrYAXjQorajncORbS0zI48LQlE2kQWg=",
        version = "v3.5.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v2",
        importpath = "go.etcd.io/etcd/client/v2",
        sum = "h1:RWfV1SX5jTU0lbCvpVQe3iPQeAHETWdOTb6pxhd77C8=",
        version = "v2.305.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_client_v3",
        importpath = "go.etcd.io/etcd/client/v3",
        sum = "h1:o0fHTNJLeO0MyVbc7I3fsCf6nrOqn5d+diSarKnB2js=",
        version = "v3.5.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_etcdctl_v3",
        importpath = "go.etcd.io/etcd/etcdctl/v3",
        sum = "h1:i8DGjR9gBRoS6NEHF3XBxxh7QwL1DyilXMCkHpyy6zM=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_etcd_etcdutl_v3",
        importpath = "go.etcd.io/etcd/etcdutl/v3",
        sum = "h1:GEAIyquWCRS0P9UAs6QmMgo36t9tT6hHNLb3g25DGNg=",
        version = "v3.5.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_pkg_v3",
        importpath = "go.etcd.io/etcd/pkg/v3",
        patch_args = ["-p1"],
        sum = "h1:st9bDWNsKkBNpP4PR1MvM/9NqUPfvYZx/YXegsYEH8M=",
        version = "v3.5.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_raft_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/raft/v3",
        sum = "h1:7r/NKAOups1YnKcfro2RvGGo2PTuizF/xh26Z2CTAzA=",
        version = "v3.5.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_server_v3",
        build_file_proto_mode = "disable",
        importpath = "go.etcd.io/etcd/server/v3",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:etcd-fix-embed-close.patch",
        ],
        sum = "h1:V6KG+yMfMSqWt+lGnhFpP5z5dRUj1BDRJ5k1fQ9DFok=",
        version = "v3.5.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_tests_v3",
        importpath = "go.etcd.io/etcd/tests/v3",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:etcd-integration-logging.patch",
        ],
        sum = "h1:vlEC8lzIyDA6Ty+vzSmnMBbrwmZ6pFQ5oTcGA+sD/Oo=",
        version = "v3.5.13",
    )
    go_repository(
        name = "io_etcd_go_etcd_v3",
        importpath = "go.etcd.io/etcd/v3",
        sum = "h1:fs7tB+L/xRDi/+p9qKuaPGCtMX6vkovLRXTqvEE98Ek=",
        version = "v3.5.0",
    )
    go_repository(
        name = "io_etcd_go_gofail",
        importpath = "go.etcd.io/gofail",
        sum = "h1:XItAMIhOojXFQMgrxjnd2EIIHun/d5qL0Pf7FzVTkFg=",
        version = "v0.1.0",
    )
    go_repository(
        name = "io_gitea_code_sdk_gitea",
        importpath = "code.gitea.io/sdk/gitea",
        sum = "h1:CdI3J82Mqn0mElyEKa5DUSr3Wi2R+qm/6uVtCkSSqSM=",
        version = "v0.11.3",
    )
    go_repository(
        name = "io_gorm_driver_mysql",
        importpath = "gorm.io/driver/mysql",
        sum = "h1:omJoilUzyrAp0xNoio88lGJCroGdIOen9hq2A/+3ifw=",
        version = "v1.0.1",
    )
    go_repository(
        name = "io_gorm_driver_postgres",
        importpath = "gorm.io/driver/postgres",
        sum = "h1:1FPESNXqIKG5JmraaH2bfCVlMQ7paLoCreFxDtqzwdc=",
        version = "v1.4.6",
    )
    go_repository(
        name = "io_gorm_driver_sqlserver",
        importpath = "gorm.io/driver/sqlserver",
        sum = "h1:nMtEeKqv2R/vv9FoHUFWfXfP6SskAgRar0TPlZV1stk=",
        version = "v1.4.2",
    )
    go_repository(
        name = "io_gorm_gorm",
        importpath = "gorm.io/gorm",
        sum = "h1:wy98aq9oFEetsc4CAbKD2SoBCdMzsbSIvSUUFJuHi5s=",
        version = "v1.24.6",
    )
    go_repository(
        name = "io_k8s_api",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/api",
        replace = "k8s.io/api",
        sum = "h1:+ZhRj+28QT4UOH+BKznu4CBgPWgkXO7XAvMcMl0qKvI=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_apiextensions_apiserver",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apiextensions-apiserver",
        replace = "k8s.io/apiextensions-apiserver",
        sum = "h1:l7Eue2t6QiLHErfn2vwK4KgF4NeDgjQkCXtEbOocKIE=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_apimachinery",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apimachinery",
        replace = "k8s.io/apimachinery",
        sum = "h1:fEMcnBj6qkzzPGSVsAZtQThU62SmQ4ZymlXRC5yFSCg=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_apiserver",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/apiserver",
        replace = "k8s.io/apiserver",
        sum = "h1:ACouHiYl1yFI2VFI3YGM+lvxgy6ir4yK2oLOsLI1/tw=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_cli_runtime",
        importpath = "k8s.io/cli-runtime",
        replace = "k8s.io/cli-runtime",
        sum = "h1:ooM40eEJusbgHNEqnHziN9ZpLN5U4WcQGsdLKVxpkKE=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_client_go",
        importpath = "k8s.io/client-go",
        replace = "k8s.io/client-go",
        sum = "h1:sBIVJdojUNPDU/jObC+18tXWcTJVcwyqS9diGdWHk50=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_cloud_provider",
        importpath = "k8s.io/cloud-provider",
        replace = "k8s.io/cloud-provider",
        sum = "h1:yov6r02v7sMUNNvzEz51LtL2krn2c1wsC+dy/8BxKQI=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_cluster_bootstrap",
        importpath = "k8s.io/cluster-bootstrap",
        replace = "k8s.io/cluster-bootstrap",
        sum = "h1:9PQ5phjWTxmPFKPEzTG6QJzPaUIfuW2RqcHDME5gqPg=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_code_generator",
        importpath = "k8s.io/code-generator",
        replace = "k8s.io/code-generator",
        sum = "h1:ZY1+aGkqZVwKIyGsOzquaeZ5rSfE6wZHur8z3jQAaiw=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_component_base",
        importpath = "k8s.io/component-base",
        replace = "k8s.io/component-base",
        sum = "h1:pqGBczYoW1sno8q9ObExUqrYSKhtE5rW3y6gX88GZII=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_component_helpers",
        importpath = "k8s.io/component-helpers",
        replace = "k8s.io/component-helpers",
        sum = "h1:kDMYLiWEYeWU7H6jBI+Ua1i2hqNh0DzqDHNIppFC3po=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_controller_manager",
        importpath = "k8s.io/controller-manager",
        replace = "k8s.io/controller-manager",
        sum = "h1:tC7V7IdGUW2I4de3bXx4m2fS3naP7VlCYlECCajK9fU=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_cri_api",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/cri-api",
        replace = "k8s.io/cri-api",
        sum = "h1:4KR5W6ziqfGzKYVmFG9AEOJzxNbCPyZMoeCeIlK9jew=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_csi_translation_lib",
        importpath = "k8s.io/csi-translation-lib",
        replace = "k8s.io/csi-translation-lib",
        sum = "h1:ZcFVMWDHg7feW3mtdl+xClgmw1Yxv7m9ysOKt8h3K8Y=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_dynamic_resource_allocation",
        importpath = "k8s.io/dynamic-resource-allocation",
        replace = "k8s.io/dynamic-resource-allocation",
        sum = "h1:wEhjNbVPymPEY5Db4UXPiQkioHV/4MHDzAkf+1TLaNM=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_endpointslice",
        importpath = "k8s.io/endpointslice",
        replace = "k8s.io/endpointslice",
        sum = "h1:d11My11frxpBD+55vx1lHeO2JlH9YQS8Xqt7EInvPLA=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_gengo",
        importpath = "k8s.io/gengo",
        sum = "h1:GohjlNKauSai7gN4wsJkeZ3WAJx4Sh+oT/b5IYn5suA=",
        version = "v0.0.0-20210813121822-485abfe95c7c",
    )
    go_repository(
        name = "io_k8s_gengo_v2",
        importpath = "k8s.io/gengo/v2",
        sum = "h1:NGrVE502P0s0/1hudf8zjgwki1X/TByhmAoILTarmzo=",
        version = "v2.0.0-20240228010128-51d4e06bde70",
    )
    go_repository(
        name = "io_k8s_klog",
        importpath = "k8s.io/klog",
        sum = "h1:Pt+yjF5aB1xDSVbau4VsWe+dQNzA0qv1LlXdC2dF6Q8=",
        version = "v1.0.0",
    )
    go_repository(
        name = "io_k8s_klog_v2",
        importpath = "k8s.io/klog/v2",
        sum = "h1:QXU6cPEOIslTGvZaXvFWiP9VKyeet3sawzTOvdXb4Vw=",
        version = "v2.120.1",
    )
    go_repository(
        name = "io_k8s_kms",
        importpath = "k8s.io/kms",
        replace = "k8s.io/kms",
        sum = "h1:VSZILO/tkzrz5Tu2j+yFQZ2Dc5JerQZX2GqhFJbQrfw=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_kube_aggregator",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kube-aggregator",
        replace = "k8s.io/kube-aggregator",
        sum = "h1:0+yk/ED6foCprY8VmkDPUhngjaAPKsNTXB/UrtvbIz0=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_kube_controller_manager",
        importpath = "k8s.io/kube-controller-manager",
        replace = "k8s.io/kube-controller-manager",
        sum = "h1:lLQKdk8iWrV1j7gKH0kirfc0YmlSl7xoHuqc8r7Fel0=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_kube_openapi",
        importpath = "k8s.io/kube-openapi",
        sum = "h1:BZqlfIlq5YbRMFko6/PM7FjZpUb45WallggurYhKGag=",
        version = "v0.0.0-20240228011516-70dd3763d340",
    )
    go_repository(
        name = "io_k8s_kube_proxy",
        importpath = "k8s.io/kube-proxy",
        replace = "k8s.io/kube-proxy",
        sum = "h1:rBLnKQzXO5u2KO7t3rqYkHf5/oG46ZjqcqfEUw4yOEs=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_kube_scheduler",
        importpath = "k8s.io/kube-scheduler",
        replace = "k8s.io/kube-scheduler",
        sum = "h1:2uR6qVkAV4ySzAQws5XDTWu62mRCxwm2m9oS0EIhAQI=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_kubectl",
        importpath = "k8s.io/kubectl",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:kubectl-fix-embed.patch",
        ],
        replace = "k8s.io/kubectl",
        sum = "h1:cgKNIvsOiufgcs4yjvgkK0+aPCfa8pUwzXdJtkbhsH8=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_kubelet",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/kubelet",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:kubelet-fix-path.patch",
        ],
        replace = "k8s.io/kubelet",
        sum = "h1:Ck4E/pHndI20IzDXxS57dElhDGASPO5pzXF7BcKfmCY=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_kubernetes",
        build_extra_args = [
            "-exclude=test/typecheck/testdata",
        ],
        build_file_proto_mode = "disable",
        build_tags = ["providerless", "dockerless"],  #keep
        importpath = "k8s.io/kubernetes",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:k8s-native-metrics.patch",
            "//third_party/go/patches:k8s-removed-block-device-pseudo-locks.patch",
            "//third_party/go/patches:k8s-add-pause-build-file.patch",
        ],
        pre_patches = [
            "//third_party/go/patches:k8s-drop-legacy-log-path.patch",
            "//third_party/go/patches:k8s-reserve-metropolis-ports.patch",
        ],
        replace = "k8s.io/kubernetes",
        sum = "h1:11WhS78OYX/lnSy6TXxPO6Hk+E5K9ZNrEsk9JgMSX8I=",
        version = "v1.30.2",
    )
    go_repository(
        name = "io_k8s_legacy_cloud_providers",
        importpath = "k8s.io/legacy-cloud-providers",
        replace = "k8s.io/legacy-cloud-providers",
        sum = "h1:RfMtmbAPvTn7+nkHRWXpGeaif4x7VBOU2SAZ2BdFEdI=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_metrics",
        build_file_proto_mode = "disable",
        importpath = "k8s.io/metrics",
        replace = "k8s.io/metrics",
        sum = "h1:zj4kIPTCfEbY0RHEogpA7QtlItU7xaO11+Gz1zVDxlc=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_mount_utils",
        importpath = "k8s.io/mount-utils",
        patch_args = ["-p1"],
        pre_patches = [
            "//third_party/go/patches:k8s-native-mounter.patch",
        ],
        replace = "k8s.io/mount-utils",
        sum = "h1:2KDVY9hXyDyRw9EO4lmox4+Nn5atVOq+4ffZ/br2aAU=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_pod_security_admission",
        importpath = "k8s.io/pod-security-admission",
        replace = "k8s.io/pod-security-admission",
        sum = "h1:UlHnkvvOr+rgQplOqD+SHzLUF8EgKIOCpDU8kaMeTQQ=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_sample_apiserver",
        importpath = "k8s.io/sample-apiserver",
        replace = "k8s.io/sample-apiserver",
        sum = "h1:bnkF2ER7MQ6mazUmNSi5niRFKooL3uU6e1RhIgZMres=",
        version = "v0.30.2",
    )
    go_repository(
        name = "io_k8s_sigs_apiserver_network_proxy_konnectivity_client",
        importpath = "sigs.k8s.io/apiserver-network-proxy/konnectivity-client",
        sum = "h1:/U5vjBbQn3RChhv7P11uhYvCSm5G2GaIi5AIGBS6r4c=",
        version = "v0.29.0",
    )
    go_repository(
        name = "io_k8s_sigs_json",
        importpath = "sigs.k8s.io/json",
        sum = "h1:EDPBXCAspyGV4jQlpZSudPeMmr1bNJefnuqLsRAsHZo=",
        version = "v0.0.0-20221116044647-bc3834ca7abd",
    )
    go_repository(
        name = "io_k8s_sigs_knftables",
        importpath = "sigs.k8s.io/knftables",
        sum = "h1:VzKQoDMCGBOH8c85sGrWSXSPCS0XrIpEfOlcCLBXiC0=",
        version = "v0.0.14",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize_api",
        importpath = "sigs.k8s.io/kustomize/api",
        sum = "h1:XX3Ajgzov2RKUdc5jW3t5jwY7Bo7dcRm+tFxT+NfgY0=",
        version = "v0.13.5-0.20230601165947-6ce0bf390ce3",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize_cmd_config",
        importpath = "sigs.k8s.io/kustomize/cmd/config",
        sum = "h1:YyoHHbxxsLUts/gWLGgIQkdT82ekp3zautbpcml54vc=",
        version = "v0.11.2",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize_kustomize_v5",
        importpath = "sigs.k8s.io/kustomize/kustomize/v5",
        sum = "h1:vq2TtoDcQomhy7OxXLUOzSbHMuMYq0Bjn93cDtJEdKw=",
        version = "v5.0.4-0.20230601165947-6ce0bf390ce3",
    )
    go_repository(
        name = "io_k8s_sigs_kustomize_kyaml",
        importpath = "sigs.k8s.io/kustomize/kyaml",
        sum = "h1:W6cLQc5pnqM7vh3b7HvGNfXrJ/xL6BDMS0v1V/HHg5U=",
        version = "v0.14.3-0.20230601165947-6ce0bf390ce3",
    )
    go_repository(
        name = "io_k8s_sigs_structured_merge_diff_v4",
        importpath = "sigs.k8s.io/structured-merge-diff/v4",
        sum = "h1:150L+0vs/8DA78h1u02ooW1/fFq/Lwr+sGiqlzvrtq4=",
        version = "v4.4.1",
    )
    go_repository(
        name = "io_k8s_sigs_yaml",
        importpath = "sigs.k8s.io/yaml",
        sum = "h1:a2VclLzOGrwOHDiV8EfBGhvjHvP46CtW5j6POvhYGGo=",
        version = "v1.3.0",
    )
    go_repository(
        name = "io_k8s_system_validators",
        importpath = "k8s.io/system-validators",
        sum = "h1:tq05tdO9zdJZnNF3SXrq6LE7Knc/KfJm5wk68467JDg=",
        version = "v1.8.0",
    )
    go_repository(
        name = "io_k8s_utils",
        importpath = "k8s.io/utils",
        sum = "h1:sgn3ZU783SCgtaSJjpcVVlRqd6GSnlTLKgpAAttJvpI=",
        version = "v0.0.0-20230726121419-3b25d923346b",
    )
    go_repository(
        name = "io_nhooyr_websocket",
        importpath = "nhooyr.io/websocket",
        sum = "h1:s+C3xAMLwGmlI31Nyn/eAehUlZPwfYZu2JXM621Q5/k=",
        version = "v1.8.6",
    )
    go_repository(
        name = "io_opencensus_go",
        importpath = "go.opencensus.io",
        sum = "h1:y73uSU6J157QMP2kn2r30vwW1A2W2WFwSCGnAVxeaD0=",
        version = "v0.24.0",
    )
    go_repository(
        name = "io_opencensus_go_contrib_exporter_aws",
        importpath = "contrib.go.opencensus.io/exporter/aws",
        sum = "h1:YsbWYxDZkC7x2OxlsDEYvvEXZ3cBI3qBgUK5BqkZvRw=",
        version = "v0.0.0-20181029163544-2befc13012d0",
    )
    go_repository(
        name = "io_opencensus_go_contrib_exporter_ocagent",
        importpath = "contrib.go.opencensus.io/exporter/ocagent",
        sum = "h1:TKXjQSRS0/cCDrP7KvkgU6SmILtF/yV2TOs/02K/WZQ=",
        version = "v0.5.0",
    )
    go_repository(
        name = "io_opencensus_go_contrib_exporter_stackdriver",
        importpath = "contrib.go.opencensus.io/exporter/stackdriver",
        sum = "h1:lIFYmQsqejvlq+GobFUbC5F0prD5gvhP6r0gWLZRDq4=",
        version = "v0.13.8",
    )
    go_repository(
        name = "io_opencensus_go_contrib_integrations_ocsql",
        importpath = "contrib.go.opencensus.io/integrations/ocsql",
        sum = "h1:kfg5Yyy1nYUrqzyfW5XX+dzMASky8IJXhtHe0KTYNS4=",
        version = "v0.1.4",
    )
    go_repository(
        name = "io_opencensus_go_contrib_resource",
        importpath = "contrib.go.opencensus.io/resource",
        sum = "h1:4r2CANuYhKGmYWP02+5E94rLRcS/YeD+KlxSrOsMxk0=",
        version = "v0.1.1",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib",
        importpath = "go.opentelemetry.io/contrib",
        sum = "h1:ubFQUn0VCZ0gPwIoJfBJVpeBlyRMxu8Mm/huKWYd9p0=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_github_com_emicklei_go_restful_otelrestful",
        importpath = "go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful",
        sum = "h1:Z6SbqeRZAl2OczfkFOqLx1BeYBDYehNjEnqluD7581Y=",
        version = "v0.42.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_google_golang_org_grpc_otelgrpc",
        importpath = "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc",
        sum = "h1:UNQQKPfTDe1J81ViolILjTKPr9WetKW6uei2hFgJmFs=",
        version = "v0.47.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_instrumentation_net_http_otelhttp",
        importpath = "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp",
        sum = "h1:sv9kVfal0MK0wBMCOGr+HeJm9v803BkJxGrk2au7j08=",
        version = "v0.47.0",
    )
    go_repository(
        name = "io_opentelemetry_go_contrib_propagators_b3",
        importpath = "go.opentelemetry.io/contrib/propagators/b3",
        sum = "h1:ImOVvHnku8jijXqkwCSyYKRDt2YrnGXD4BbhcpfbfJo=",
        version = "v1.17.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel",
        importpath = "go.opentelemetry.io/otel",
        sum = "h1:xS7Ku+7yTFvDfDraDIJVpw7XPyuHlB9MCiqqX5mcJ6Y=",
        version = "v1.22.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp",
        importpath = "go.opentelemetry.io/otel/exporters/otlp",
        sum = "h1:PTNgq9MRmQqqJY0REVbZFvwkYOA85vbdQU/nVfxDyqg=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_internal_retry",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/internal/retry",
        sum = "h1:R/OBkMoGgfy2fLhs2QhkCI1w4HLEQX92GCcJB6SSdNk=",
        version = "v1.3.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace",
        build_directives = [
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal @io_opentelemetry_go_otel//exporters/otlp/internal",
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal/envconfig @io_opentelemetry_go_otel//exporters/otlp/internal/envconfig",
        ],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace",
        sum = "h1:DeFD0VgTZ+Cj6hxravYYZE2W4GlneVH81iAOPjZkzk8=",
        version = "v1.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracegrpc",
        build_directives = [
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal @io_opentelemetry_go_otel//exporters/otlp/internal",
        ],
        build_file_proto_mode = "disable",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc",
        sum = "h1:gvmNvqrPYovvyRmCSygkUDyL8lC5Tl845MLEwqpxhEU=",
        version = "v1.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_exporters_otlp_otlptrace_otlptracehttp",
        build_directives = [
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/internal @io_opentelemetry_go_otel//exporters/otlp/internal",
            "gazelle:resolve go go.opentelemetry.io/otel/exporters/otlp/otlptrace/internal/retry @io_opentelemetry_go_otel_exporters_otlp_internal_retry",
        ],
        build_file_proto_mode = "disable_global",
        importpath = "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp",
        sum = "h1:IeMeyr1aBvBiPVYihXIaeIZba6b8E1bYp7lbdxK8CQg=",
        version = "v1.19.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_metric",
        importpath = "go.opentelemetry.io/otel/metric",
        sum = "h1:lypMQnGyJYeuYPhOM/bgjbFM6WE44W1/T45er4d8Hhg=",
        version = "v1.22.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_oteltest",
        importpath = "go.opentelemetry.io/otel/oteltest",
        sum = "h1:HiITxCawalo5vQzdHfKeZurV8x7ljcqAgiWzF6Vaeaw=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk",
        importpath = "go.opentelemetry.io/otel/sdk",
        sum = "h1:FTt8qirL1EysG6sTQRZ5TokkU8d0ugCj8htOgThZXQ8=",
        version = "v1.21.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_export_metric",
        importpath = "go.opentelemetry.io/otel/sdk/export/metric",
        sum = "h1:c5VRjxCXdQlx1HjzwGdQHzZaVI82b5EbBgOu2ljD92g=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_sdk_metric",
        importpath = "go.opentelemetry.io/otel/sdk/metric",
        sum = "h1:7ao1wpzHRVKf0OQ7GIxiQJA6X7DLX9o14gmVon7mMK8=",
        version = "v0.20.0",
    )
    go_repository(
        name = "io_opentelemetry_go_otel_trace",
        importpath = "go.opentelemetry.io/otel/trace",
        sum = "h1:Hg6pPujv0XG9QaVbGOBVHunyuLcCC3jN7WEhPx83XD0=",
        version = "v1.22.0",
    )
    go_repository(
        name = "io_opentelemetry_go_proto_otlp",
        importpath = "go.opentelemetry.io/proto/otlp",
        sum = "h1:T0TX0tmXU8a3CbNXzEKGeU5mIVOdf0oykP+u2lIVU/I=",
        version = "v1.0.0",
    )
    go_repository(
        name = "io_rsc_binaryregexp",
        importpath = "rsc.io/binaryregexp",
        sum = "h1:HfqmD5MEmC0zvwBuF187nq9mdnXjXsSivRiXN7SmRkE=",
        version = "v0.2.0",
    )
    go_repository(
        name = "io_rsc_pdf",
        importpath = "rsc.io/pdf",
        sum = "h1:k1MczvYDUvJBe93bYd7wrZLLUEcLZAuF824/I4e5Xr4=",
        version = "v0.1.1",
    )
    go_repository(
        name = "io_rsc_quote_v3",
        importpath = "rsc.io/quote/v3",
        sum = "h1:9JKUTTIUgS6kzR9mK1YuGKv6Nl+DijDNIc0ghT58FaY=",
        version = "v3.1.0",
    )
    go_repository(
        name = "io_rsc_sampler",
        importpath = "rsc.io/sampler",
        sum = "h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=",
        version = "v1.3.0",
    )
    go_repository(
        name = "net_howett_plist",
        importpath = "howett.net/plist",
        sum = "h1:7CrbWYbPPO/PyNy38b2EB/+gYbjCe2DXBxgtOOZbSQM=",
        version = "v1.0.0",
    )
    go_repository(
        name = "net_starlark_go",
        importpath = "go.starlark.net",
        sum = "h1:VdD38733bfYv5tUZwEIskMM93VanwNIi5bIKnDrJdEY=",
        version = "v0.0.0-20230525235612-a134d8f9ddca",
    )
    go_repository(
        name = "org_bazil_fuse",
        importpath = "bazil.org/fuse",
        sum = "h1:SRsZGA7aFnCZETmov57jwPrWuTmaZK6+4R4v5FUe1/c=",
        version = "v0.0.0-20200407214033-5883e5a4b512",
    )
    go_repository(
        name = "org_bitbucket_bertimus9_systemstat",
        importpath = "bitbucket.org/bertimus9/systemstat",
        sum = "h1:n0aLnh2Jo4nBUBym9cE5PJDG8GT6g+4VuS2Ya2jYYpA=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_bitbucket_creachadair_shell",
        importpath = "bitbucket.org/creachadair/shell",
        sum = "h1:reJflDbKqnlnqb4Oo2pQ1/BqmY/eCWcNGHrIUO8qIzc=",
        version = "v0.0.6",
    )
    go_repository(
        name = "org_gioui",
        importpath = "gioui.org",
        sum = "h1:K72hopUosKG3ntOPNG4OzzbuhxGuVf06fa2la1/H/Ho=",
        version = "v0.0.0-20210308172011-57750fc8a0a6",
    )
    go_repository(
        name = "org_go4_intern",
        importpath = "go4.org/intern",
        sum = "h1:UXLjNohABv4S58tHmeuIZDO6e3mHpW2Dx33gaNt03LE=",
        version = "v0.0.0-20211027215823-ae77deb06f29",
    )
    go_repository(
        name = "org_go4_unsafe_assume_no_moving_gc",
        importpath = "go4.org/unsafe/assume-no-moving-gc",
        sum = "h1:FyBZqvoA/jbNzuAWLQE2kG820zMAkcilx6BMjGbL/E4=",
        version = "v0.0.0-20220617031537-928513b29760",
    )
    go_repository(
        name = "org_golang_google_api",
        importpath = "google.golang.org/api",
        sum = "h1:Vhs54HkaEpkMBdgGdOT2P6F0csGG/vxDS0hWHJzmmps=",
        version = "v0.162.0",
    )
    go_repository(
        name = "org_golang_google_appengine",
        importpath = "google.golang.org/appengine",
        sum = "h1:IhEN5q69dyKagZPYMSdIjS2HqprW324FRQZJcGqPAsM=",
        version = "v1.6.8",
    )
    go_repository(
        name = "org_golang_google_cloud",
        importpath = "google.golang.org/cloud",
        sum = "h1:Cpp2P6TPjujNoC5M2KHY6g7wfyLYfIWRZaSdIKfDasA=",
        version = "v0.0.0-20151119220103-975617b05ea8",
    )
    go_repository(
        name = "org_golang_google_genproto",
        importpath = "google.golang.org/genproto",
        sum = "h1:F6qOa9AZTYJXOUEr4jDysRDLrm4PHePlge4v4TGAlxY=",
        version = "v0.0.0-20240227224415-6ceb2ff114de",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_api",
        build_file_proto_mode = "disable_global",
        importpath = "google.golang.org/genproto/googleapis/api",
        sum = "h1:Zz7rLWqp0ApfsR/l7+zSHhY3PMiH2xqgxlfYfAfNpoU=",
        version = "v0.0.0-20240415180920-8c6c420018be",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_bytestream",
        importpath = "google.golang.org/genproto/googleapis/bytestream",
        sum = "h1:weYsP+dNijSQVoLAb5bpUos3ciBpNU/NEVlHFKrk8pg=",
        version = "v0.0.0-20240125205218-1f4bbc51befe",
    )
    go_repository(
        name = "org_golang_google_genproto_googleapis_rpc",
        build_file_proto_mode = "disable_global",
        importpath = "google.golang.org/genproto/googleapis/rpc",
        sum = "h1:LG9vZxsWGOmUKieR8wPAUR3u3MpnYFQZROPIMaXh7/A=",
        version = "v0.0.0-20240415180920-8c6c420018be",
    )
    go_repository(
        name = "org_golang_google_grpc",
        build_file_proto_mode = "disable",
        importpath = "google.golang.org/grpc",
        sum = "h1:MUeiw1B2maTVZthpU5xvASfTh3LDbxHd6IJ6QQVU+xM=",
        version = "v1.63.2",
    )
    go_repository(
        name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
        importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
        sum = "h1:M1YKkFIboKNieVO5DLUEVzQfGwJD30Nv2jfUgzb5UcE=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_golang_google_protobuf",
        build_directives = [
            "gazelle:proto disable",  # https://github.com/bazelbuild/rules_go/issues/3906
        ],
        build_extra_args = [
            "-exclude=**/testdata",
        ],
        importpath = "google.golang.org/protobuf",
        sum = "h1:uNO2rsAINq/JlFpSdYEKIZ0uKD/R9cpdv0T+yoGwGmI=",
        version = "v1.33.0",
    )
    go_repository(
        name = "org_golang_x_arch",
        importpath = "golang.org/x/arch",
        sum = "h1:02VY4/ZcO/gBOH6PUaoiptASxtXU10jazRCP865E97k=",
        version = "v0.3.0",
    )
    go_repository(
        name = "org_golang_x_crypto",
        importpath = "golang.org/x/crypto",
        sum = "h1:X31++rzVUdKhX5sWmSOFZxx8UW/ldWx55cbf08iNAMA=",
        version = "v0.21.0",
    )
    go_repository(
        name = "org_golang_x_exp",
        importpath = "golang.org/x/exp",
        replace = "golang.org/x/exp",
        sum = "h1:FRnLl4eNAQl8hwxVVC17teOw8kdjVDVAiFMtgUdTSRQ=",
        version = "v0.0.0-20231110203233-9a3e6036ecaa",
    )
    go_repository(
        name = "org_golang_x_exp_typeparams",
        importpath = "golang.org/x/exp/typeparams",
        sum = "h1:Jw5wfR+h9mnIYH+OtGT2im5wV1YGGDora5vTv/aa5bE=",
        version = "v0.0.0-20221208152030-732eee02a75a",
    )
    go_repository(
        name = "org_golang_x_image",
        importpath = "golang.org/x/image",
        sum = "h1:TcHcE0vrmgzNH1v3ppjcMGbhG5+9fMuvOmUYwNEF4q4=",
        version = "v0.0.0-20220302094943-723b81ca9867",
    )
    go_repository(
        name = "org_golang_x_lint",
        importpath = "golang.org/x/lint",
        sum = "h1:VLliZ0d+/avPrXXH+OakdXhpJuEoBZuwh1m2j7U6Iug=",
        version = "v0.0.0-20210508222113-6edffad5e616",
    )
    go_repository(
        name = "org_golang_x_mod",
        build_extra_args = [
            "-go_naming_convention=import_alias",
            "-go_naming_convention_external=import_alias",
        ],
        importpath = "golang.org/x/mod",
        sum = "h1:SernR4v+D55NyBH2QiEQrlBAnj1ECL6AGrA5+dPaMY8=",
        version = "v0.15.0",
    )
    go_repository(
        name = "org_golang_x_net",
        importpath = "golang.org/x/net",
        sum = "h1:7EYJ93RZ9vYSZAIb2x3lnuvqO5zneoD6IvWjuhfxjTs=",
        version = "v0.23.0",
    )
    go_repository(
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        sum = "h1:6m3ZPmLEFdVxKKWnKq4VqZ60gutO35zm+zrAHVmHyDQ=",
        version = "v0.17.0",
    )
    go_repository(
        name = "org_golang_x_sync",
        build_extra_args = [
            "-go_naming_convention=import_alias",
            "-go_naming_convention_external=import_alias",
        ],
        importpath = "golang.org/x/sync",
        sum = "h1:5BMeUDZ7vkXGfEr1x9B4bRcTH4lpkTkpdh0T/J+qjbQ=",
        version = "v0.6.0",
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        sum = "h1:DBdB3niSjOA/O0blCZBqDefyWNYveAYMNF1Wum0DYQ4=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_telemetry",
        importpath = "golang.org/x/telemetry",
        sum = "h1:+Kc94D8UVEVxJnLXp/+FMfqQARZtWHfVrcRtcG8aT3g=",
        version = "v0.0.0-20240208230135-b75ee8823808",
    )
    go_repository(
        name = "org_golang_x_term",
        importpath = "golang.org/x/term",
        sum = "h1:FcHjZXDMxI8mM3nwhX9HlKop4C0YQvCVCdwYl2wOtE8=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_text",
        importpath = "golang.org/x/text",
        sum = "h1:ScX5w1eTa3QqT8oi6+ziP7dTV1S2+ALU0bI+0zXKWiQ=",
        version = "v0.14.0",
    )
    go_repository(
        name = "org_golang_x_time",
        importpath = "golang.org/x/time",
        sum = "h1:o7cqy6amK/52YcAKIPlM3a+Fpj35zvRj2TP+e1xFSfk=",
        version = "v0.5.0",
    )
    go_repository(
        name = "org_golang_x_tools",
        importpath = "golang.org/x/tools",
        patch_args = ["-p1"],
        patches = [
            "//third_party/go/patches:goimports-group-merging.patch",
        ],
        sum = "h1:k8NLag8AGHnn+PHbl7g43CtqZAwG60vZkLqgyZgIHgQ=",
        version = "v0.18.0",
    )
    go_repository(
        name = "org_golang_x_xerrors",
        importpath = "golang.org/x/xerrors",
        sum = "h1:H2TDz8ibqkAF6YGhCdN3jS9O0/s90v0rJh3X/OLHEUk=",
        version = "v0.0.0-20220907171357-04be3eba64a2",
    )
    go_repository(
        name = "org_gonum_v1_gonum",
        importpath = "gonum.org/v1/gonum",
        sum = "h1:f1IJhK4Km5tBJmaiJXtk/PkL4cdVX6J+tGiM187uT5E=",
        version = "v0.11.0",
    )
    go_repository(
        name = "org_gonum_v1_netlib",
        importpath = "gonum.org/v1/netlib",
        sum = "h1:OE9mWmgKkjJyEmDAAtGMPjXu+YNeGvK9VTSHY6+Qihc=",
        version = "v0.0.0-20190313105609-8cb42192e0e0",
    )
    go_repository(
        name = "org_gonum_v1_plot",
        importpath = "gonum.org/v1/plot",
        sum = "h1:dnifSs43YJuNMDzB7v8wV64O4ABBHReuAVAoBxqBqS4=",
        version = "v0.10.1",
    )
    go_repository(
        name = "org_modernc_b",
        importpath = "modernc.org/b",
        sum = "h1:vpvqeyp17ddcQWF29Czawql4lDdABCDRbXRAS4+aF2o=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_cc_v3",
        importpath = "modernc.org/cc/v3",
        sum = "h1:P3g79IUS/93SYhtoeaHW+kRCIrYaxJ27MFPv+7kaTOw=",
        version = "v3.40.0",
    )
    go_repository(
        name = "org_modernc_ccgo_v3",
        importpath = "modernc.org/ccgo/v3",
        sum = "h1:Mkgdzl46i5F/CNR/Kj80Ri59hC8TKAhZrYSaqvkwzUw=",
        version = "v3.16.13",
    )
    go_repository(
        name = "org_modernc_ccorpus",
        importpath = "modernc.org/ccorpus",
        sum = "h1:J16RXiiqiCgua6+ZvQot4yUuUy8zxgqbqEEUuGPlISk=",
        version = "v1.11.6",
    )
    go_repository(
        name = "org_modernc_db",
        importpath = "modernc.org/db",
        sum = "h1:2c6NdCfaLnshSvY7OU09cyAY0gYXUZj4lmg5ItHyucg=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_file",
        importpath = "modernc.org/file",
        sum = "h1:9/PdvjVxd5+LcWUQIfapAWRGOkDLK90rloa8s/au06A=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_fileutil",
        importpath = "modernc.org/fileutil",
        sum = "h1:Z1AFLZwl6BO8A5NldQg/xTSjGLetp+1Ubvl4alfGx8w=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_golex",
        importpath = "modernc.org/golex",
        sum = "h1:dmSaksHMd+y6NkBsRsCShNPRaSNCNH+abrVm5/gZic8=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_modernc_httpfs",
        importpath = "modernc.org/httpfs",
        sum = "h1:AAgIpFZRXuYnkjftxTAZwMIiwEqAfk8aVB2/oA6nAeM=",
        version = "v1.0.6",
    )
    go_repository(
        name = "org_modernc_internal",
        importpath = "modernc.org/internal",
        sum = "h1:XMDsFDcBDsibbBnHB2xzljZ+B1yrOVLEFkKL2u15Glw=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_libc",
        importpath = "modernc.org/libc",
        sum = "h1:4U7v51GyhlWqQmwCHj28Rdq2Yzwk55ovjFrdPjs8Hb0=",
        version = "v1.22.2",
    )
    go_repository(
        name = "org_modernc_lldb",
        importpath = "modernc.org/lldb",
        sum = "h1:6vjDJxQEfhlOLwl4bhpwIz00uyFK4EmSYcbwqwbynsc=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_mathutil",
        importpath = "modernc.org/mathutil",
        sum = "h1:fRe9+AmYlaej+64JsEEhoWuAYBkOtQiMEU7n/XgfYi4=",
        version = "v1.6.0",
    )
    go_repository(
        name = "org_modernc_memory",
        importpath = "modernc.org/memory",
        sum = "h1:N+/8c5rE6EqugZwHii4IFsaJ7MUhoWX07J5tC/iI5Ds=",
        version = "v1.5.0",
    )
    go_repository(
        name = "org_modernc_opt",
        importpath = "modernc.org/opt",
        sum = "h1:3XOZf2yznlhC+ibLltsDGzABUGVx8J6pnFMS3E4dcq4=",
        version = "v0.1.3",
    )
    go_repository(
        name = "org_modernc_parser",
        importpath = "modernc.org/parser",
        sum = "h1:XoClYpoz2xHEDIteSQ7tICOTFcNwBI7XRCeghUS6SNI=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_modernc_ql",
        importpath = "modernc.org/ql",
        sum = "h1:bIQ/trWNVjQPlinI6jdOQsi195SIturGo3mp5hsDqVU=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_modernc_sortutil",
        importpath = "modernc.org/sortutil",
        sum = "h1:VQGxbQGcHaQeB/BX9TQjrHFmOA0bounO1X/jvOfRo6Q=",
        version = "v1.1.1",
    )
    go_repository(
        name = "org_modernc_sqlite",
        importpath = "modernc.org/sqlite",
        sum = "h1:S2uFiaNPd/vTAP/4EmyY8Qe2Quzu26A2L1e25xRNTio=",
        version = "v1.18.2",
    )
    go_repository(
        name = "org_modernc_strutil",
        importpath = "modernc.org/strutil",
        sum = "h1:agBi9dp1I+eOnxXeiZawM8F4LawKv4NzGWSaLfyeNZA=",
        version = "v1.2.0",
    )
    go_repository(
        name = "org_modernc_tcl",
        importpath = "modernc.org/tcl",
        sum = "h1:5PQgL/29XkQ9wsEmmNPjzKs+7iPCaYqUJAhzPvQbjDA=",
        version = "v1.13.2",
    )
    go_repository(
        name = "org_modernc_token",
        importpath = "modernc.org/token",
        sum = "h1:Xl7Ap9dKaEs5kLoOQeQmPWevfnk/DM5qcLcYlA8ys6Y=",
        version = "v1.1.0",
    )
    go_repository(
        name = "org_modernc_y",
        importpath = "modernc.org/y",
        sum = "h1:U3EAg4VQmj2eoAUnMFcv+KXxVQFT19ZIA1mO1XX0b1s=",
        version = "v1.0.9",
    )
    go_repository(
        name = "org_modernc_z",
        importpath = "modernc.org/z",
        sum = "h1:RTNHdsrOpeoSeOF4FbzTo8gBYByaJ5xT7NgZ9ZqRiJM=",
        version = "v1.5.1",
    )
    go_repository(
        name = "org_modernc_zappy",
        importpath = "modernc.org/zappy",
        sum = "h1:dPVaP+3ueIUv4guk8PuZ2wiUGcJ1WUVvIheeSSTD0yk=",
        version = "v1.0.0",
    )
    go_repository(
        name = "org_mongodb_go_mongo_driver",
        importpath = "go.mongodb.org/mongo-driver",
        sum = "h1:ny3p0reEpgsR2cfA5cjgwFZg3Cv/ofFh/8jbhGtz9VI=",
        version = "v1.7.5",
    )
    go_repository(
        name = "org_mozilla_go_pkcs7",
        importpath = "go.mozilla.org/pkcs7",
        sum = "h1:A/5uWzF44DlIgdm/PQFwfMkW0JX+cIcQi/SwLAmZP5M=",
        version = "v0.0.0-20200128120323-432b2356ecb1",
    )
    go_repository(
        name = "org_uber_go_atomic",
        importpath = "go.uber.org/atomic",
        sum = "h1:ZvwS0R+56ePWxUNi+Atn9dWONBPp/AUETXlHW0DxSjE=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_uber_go_goleak",
        importpath = "go.uber.org/goleak",
        sum = "h1:2K3zAYmnTNqV73imy9J1T3WC+gmCePx2hEGkimedGto=",
        version = "v1.3.0",
    )
    go_repository(
        name = "org_uber_go_multierr",
        importpath = "go.uber.org/multierr",
        sum = "h1:blXXJkSxSSfBVBlC76pxqeO+LN3aDfLQo+309xJstO0=",
        version = "v1.11.0",
    )
    go_repository(
        name = "org_uber_go_tools",
        importpath = "go.uber.org/tools",
        sum = "h1:0mgffUl7nfd+FpvXMVz4IDEaUSmT1ysygQC7qYo7sG4=",
        version = "v0.0.0-20190618225709-2cfd321de3ee",
    )
    go_repository(
        name = "org_uber_go_zap",
        importpath = "go.uber.org/zap",
        sum = "h1:sI7k6L95XOKS281NhVKOFCUNIvv9e0w4BF8N3u+tCRo=",
        version = "v1.26.0",
    )
    go_repository(
        name = "tools_gotest",
        importpath = "gotest.tools",
        sum = "h1:VsBPFP1AI068pPrMxtb/S8Zkgf9xEmTLJjfM+P5UIEo=",
        version = "v2.2.0+incompatible",
    )
    go_repository(
        name = "tools_gotest_v3",
        importpath = "gotest.tools/v3",
        sum = "h1:ZazjZUfuVeZGLAmlKKuyv3IKP5orXcwtOwDQH6YVr6o=",
        version = "v3.4.0",
    )
