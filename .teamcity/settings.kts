import jetbrains.buildServer.configs.kotlin.*
import jetbrains.buildServer.configs.kotlin.buildSteps.nodeJS
import jetbrains.buildServer.configs.kotlin.triggers.vcs
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot

/*
The settings script is an entry point for defining a TeamCity
project hierarchy. The script should contain a single call to the
project() function with a Project instance or an init function as
an argument.

VcsRoots, BuildTypes, Templates, and subprojects can be
registered inside the project using the vcsRoot(), buildType(),
template(), and subProject() methods respectively.

To debug settings scripts in command-line, run the

    mvnDebug org.jetbrains.teamcity:teamcity-configs-maven-plugin:generate

command and attach your debugger to the port 8000.

To debug in IntelliJ Idea, open the 'Maven Projects' tool window (View
-> Tool Windows -> Maven Projects), find the generate task node
(Plugins -> teamcity-configs -> teamcity-configs:generate), the
'Debug' option is available in the context menu for the task.
*/

version = "2022.04"

project {

    vcsRoot(HttpsGithubComMsugakovStackroxRefsHeadsMaster)

    buildType(Build)
}

object Build : BuildType({
    name = "Build"

    vcs {
        root(HttpsGithubComMsugakovStackroxRefsHeadsMaster)
    }

    steps {
        nodeJS {
            workingDir = "ui"
            shellScript = "yarn install --frozen-lockfile"
        }
        nodeJS {
            workingDir = "ui"
            shellScript = "yarn run lint"
        }
    }

    triggers {
        vcs {
        }
    }
})

object HttpsGithubComMsugakovStackroxRefsHeadsMaster : GitVcsRoot({
    name = "https://github.com/msugakov/stackrox#refs/heads/master"
    url = "https://github.com/msugakov/stackrox"
    branch = "refs/heads/master"
    branchSpec = "refs/heads/*"
    authMethod = password {
        userName = "git"
        password = "cks2fea026727ec59cb13bf63d1bb4e91e6u7AkF89q+MGn/gMMy4wwJg4fKgiEvufj0i5KRkbxDScsRJZnZEwzRc5y38fEpR28"
    }
    param("oauthProviderId", "tc-cloud-github-connection")
})
