node {
   stage 'git clone'
   git 'https://github.com/itouri/monGO.git'
 
   stage 'composer install'
   sh 'composer install'
 
   stage 'phpunit'
   sh 'phpunit -c phpunit.xml.dist'
}