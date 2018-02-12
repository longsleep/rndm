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
		GOBIN = '/usr/local/bin'
		GOPATH = '/workspace'
		PACKAGE = 'stash.kopano.io/kgol/rndm'
	}
	stages {
		stage('Bootstrap') {
			steps {
				sh 'go env'
				sh 'export'
				echo 'Bootstrapping..'
				sh 'mkdir -p \$GOPATH/src/\$PACKAGE && rmdir \$GOPATH/src/\$PACKAGE && ln -sv \$WORKSPACE \$GOPATH/src/\$PACKAGE'
				sh 'curl -sSL https://github.com/Masterminds/glide/releases/download/$GLIDE_VERSION/glide-$GLIDE_VERSION-linux-amd64.tar.gz | tar -vxz -C /usr/local/bin --strip=1'
				sh 'go get -v github.com/golang/lint/golint'
				sh 'go get -v github.com/tebeka/go2xunit'
				sh 'go get -v github.com/axw/gocov/...'
				sh 'go get -v github.com/AlekSi/gocov-xml'
			}
		}
		stage('Lint') {
			steps {
				echo 'Linting..'
				sh 'cd \$GOPATH/src/\$PACKAGE && golint \$(glide nv) | tee golint.txt || true'
				sh 'cd \$GOPATH/src/\$PACKAGE && go vet \$(glide nv) | tee govet.txt || true'
			}
		}
		stage('Test') {
			steps {
				echo 'Testing..'
				sh 'cd \$GOPATH/src/\$PACKAGE && go test -v -covermode=atomic -coverprofile=coverage.out \$(glide nv) | tee tests.output'
				sh 'cd \$GOPATH/src/\$PACKAGE && go2xunit -fail -input tests.output -output tests.xml'
			}
		}
		stage('Coverage') {
			steps {
				echo 'Coverage..'
				sh 'mkdir -p ./test/reports'
				sh 'cd \$GOPATH/src/\$PACKAGE && go tool cover -html=coverage.out -o test/reports/coverage.html'
				sh 'cd \$GOPATH/src/\$PACKAGE && gocov convert coverage.out | gocov-xml > coverage.xml'
			}
		}
	}
	post {
		always {
			junit allowEmptyResults: true, testResults: 'tests.xml'
			warnings parserConfigurations: [[parserName: 'Go Lint', pattern: 'golint.txt'], [parserName: 'Go Vet', pattern: 'govet.txt']], unstableTotalAll: '0'
			publishHTML([allowMissing: true, alwaysLinkToLastBuild: true, keepAll: true, reportDir: 'test/reports', reportFiles: 'coverage.html', reportName: 'Go Coverage Report HTML', reportTitles: ''])
			step([$class: 'CoberturaPublisher', autoUpdateHealth: false, autoUpdateStability: false, coberturaReportFile: 'coverage.xml', failUnhealthy: false, failUnstable: false, maxNumberOfBuilds: 0, onlyStable: false, sourceEncoding: 'ASCII', zoomCoverageChart: false])
			cleanWs()
		}
	}
}
