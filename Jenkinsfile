def imageName = 'sakkoumhamza/books-parser'
def registry = 'https://index.docker.io/v1/'

node('workers'){
    stage('Checkout'){
        checkout scm
    }

    def imageTest = docker.build("${imageName}-test", "-f Dockerfile.test .")

    stage('Pre-integration Tests'){
        parallel(
            'Quality Tests': {
                imageTest.inside {
                    sh 'golint ./...'
                }
            },
            'Unit Tests': {
                imageTest.inside {
                    sh 'go clean -testcache && GOCACHE=off go test ./...'
                }
            },
            'Security Tests': {
                withCredentials([usernamePassword(credentialsId: 'OSS', usernameVariable: 'OSS_USERNAME', passwordVariable: 'OSS_TOKEN')]) {
                    imageTest.inside('-u root:root') {
                        sh ' go list -json -m all | nancy sleuth --username $OSS_USERNAME --token $OSS_TOKEN'
                    }
            }
            }
        )
    }

    stage('Build'){
        docker.build("${imageName}:${commitID()}")
    }

    stage('Push'){
        withCredentials([usernamePassword(credentialsId: 'registry', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
            sh "docker login -u $DOCKER_USER -p $DOCKER_PASS $registry"
            docker.image("${imageName}:${commitID()}").push()
            if (env.BRANCH_NAME == 'develop') {
                docker.image("${imageName}:${commitID()}").push('develop')
            }
        }
    }
}

def commitID() {
    sh 'git rev-parse HEAD > .git/commitID'
    def commitID = readFile('.git/commitID').trim()
    sh 'rm .git/commitID'
    return commitID
}
