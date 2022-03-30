// This is a 'Jenkinsfile'-style declarative 'Pipeline' definition. It is
// executed by Jenkins for presubmit checks, ie. checks that run against an
// open Gerrit change request.

pipeline {
    agent none
    options {
        disableConcurrentBuilds()
    }
    stages {
        stage('Parallel') {
            parallel {
                stage('Test') {
                    agent {
                        node {
                            label ""
                            customWorkspace '/home/ci/monogon'
                        }
                    }
                    steps {
                        gerritCheck checks: ['jenkins:test': 'RUNNING'], message: "Running on ${env.NODE_NAME}"
                        echo "Gerrit change: ${GERRIT_CHANGE_URL}"
                        sh "git clean -fdx -e '/bazel-*'"
                        sh "JENKINS_NODE_COOKIE=dontKillMe bazel test //..."
                        sh "JENKINS_NODE_COOKIE=dontKillMe bazel test -c dbg //..."
                    }
                    post {
                        success {
                            gerritCheck checks: ['jenkins:test': 'SUCCESSFUL']
                        }
                        unsuccessful {
                            gerritCheck checks: ['jenkins:test': 'FAILED']
                        }
                    }
                }

                stage('Gazelle') {
                    agent {
                        node {
                            label ""
                            customWorkspace '/home/ci/monogon'
                        }
                    }
                    steps {
                        gerritCheck checks: ['jenkins:gazelle': 'RUNNING'], message: "Running on ${env.NODE_NAME}"
                        echo "Gerrit change: ${GERRIT_CHANGE_URL}"
                        sh "git clean -fdx -e '/bazel-*'"
                        sh "JENKINS_NODE_COOKIE=dontKillMe bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=third_party/go/repositories.bzl%go_repositories -prune"
                        sh "JENKINS_NODE_COOKIE=dontKillMe bazel run //:gazelle -- update"

                        script {
                            def diff = sh script: "git status --porcelain", returnStdout: true
                            if (diff.trim() != "") {
                                sh "git diff HEAD"
                                error """
                                    Unclean working directory after running gazelle.
                                    Please run:

                                       \$ bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=third_party/go/repositories.bzl%go_repositories -prune
                                       \$ bazel run //:gazelle -- update

                                    In your git checkout and amend the resulting diff to this changelist.
                                """
                            }
                        }
                    }
                    post {
                        success {
                            gerritCheck checks: ['jenkins:gazelle': 'SUCCESSFUL']
                        }
                        unsuccessful {
                            gerritCheck checks: ['jenkins:gazelle': 'FAILED']
                        }
                    }
                }
            }
        }
    }
}
