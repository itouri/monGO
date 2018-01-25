node {
   stage 'git clone'
   git 'https://github.com/itouri/monGO.git'
 
   stage 'unit test'
   sh 'chmod 755 ./ci/unit-tests/unit-tests.sh'
   sh './ci/unit-tests/unit-tests.sh'
 
   stage 'acceptance test'
   sh 'echo "not implement yet"'
}
