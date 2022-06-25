pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy To Staging') {
            steps {
                echo 'Deploying To Staging....'
            }
        }
        stage('Deploy To Prod') {
            // For manual production deployment
            input {
                message "Shall we deploy to production?"
                ok "Yes, Please!"
            }

            steps {
                echo 'Deploying To Prod....'
            }
        }
    }
}