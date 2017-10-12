#!/usr/bin/env groovy

pipeline {
	agent {
		docker {
			image 'golang:1.8'
			args '-u 0'
		 }
	}
	environment {
		GLIDE_VERSION = 'v0.13.0'
		GLIDE_HOME = '/tmp/.glide'
	}
	stages {
		stage('Bootstrap') {
			steps {
				echo 'Bootstrapping..'
				sh 'curl -sSL https://github.com/Masterminds/glide/releases/download/$GLIDE_VERSION/glide-$GLIDE_VERSION-linux-amd64.tar.gz | tar -vxz -C /usr/local/bin --strip=1'
				sh 'go get -v github.com/golang/lint/golint'
				sh 'go get -v github.com/tebeka/go2xunit'
			}
		}
		stage('Lint') {
			steps {
				echo 'Linting..'
				sh 'golint \$(glide nv) | tee golint.txt || true'
				sh 'go vet \$(glide nv) | tee govet.txt || true'
			}
		}
		stage('Test') {
			steps {
				echo 'Testing..'
				sh 'go test -v \$(glide nv) | tee tests.output'
				sh 'go2xunit -fail -input tests.output -output tests.xml'
			}
		}
	}
	post {
		always {
			junit 'tests.xml'
			warnings parserConfigurations: [[parserName: 'Go Lint', pattern: 'golint.txt'], [parserName: 'Go Vet', pattern: 'govet.txt']], unstableTotalAll: '0'
			cleanWs()
		}
	}
}
