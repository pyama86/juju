set -eu
# For most jobs, this is localhost, so provide it.
: ${LOCAL_JENKINS_URL=http://localhost:8080}
export JUJU_HOME=$HOME/cloud-city
if [ "$ENV" = "manual" ]; then
  source $HOME/cloud-city/ec2rc
fi


prepare_manual(){
    export INSTANCE_TYPE=m1.large
    export AMI_IMAGE=ami-bd6d40d4
    machine_0_id=$($SCRIPTS/ec2-run-instance-get-id -g manual-juju-test)
    machine_1_id=$($SCRIPTS/ec2-run-instance-get-id -g manual-juju-test)
    machine_2_id=$($SCRIPTS/ec2-run-instance-get-id -g manual-juju-test)
    $SCRIPTS/ec2-tag-job-instances $machine_0_id $machine_1_id $machine_2_id
    machine_0_name=$($SCRIPTS/ec2-get-name $machine_0_id)
    machine_1_name=$($SCRIPTS/ec2-get-name $machine_1_id)
    machine_2_name=$($SCRIPTS/ec2-get-name $machine_2_id)
    export BOOTSTRAP_HOST=$machine_0_name

    rm -f $JUJU_HOME/environments/manual.jenv

    for machine in $machine_0_name $machine_1_name $machine_2_name; do
      $SCRIPTS/wait-for-port $machine 22
    done
    DEPLOY_ARGS="--machine ssh:$machine_1_name --machine ssh:$machine_2_name"
    export MACHINES="$machine_1_name $machine_2_name"
}

artifacts_path=$WORKSPACE/artifacts
export MACHINES=""
set -x
rm * -rf
mkdir -p $artifacts_path
touch $artifacts_path/empty
afact='lastSuccessfulBuild/artifact'
wget -q $LOCAL_JENKINS_URL/job/publish-revision/$afact/new-precise.deb
# Determine BRANCH and REVNO
wget -q $LOCAL_JENKINS_URL/job/build-revision/$afact/buildvars.bash
source buildvars.bash
echo "Testing $BRANCH $REVNO on $ENV"
dpkg-deb -x $WORKSPACE/new-precise.deb extracted-bin
JUJU_BIN=$(readlink -f $(dirname $(find extracted-bin -name juju)))
export NEW_PATH=$JUJU_BIN:$PATH
if [ "$ENV" == "manual" ]; then
  $SCRIPTS/ec2-terminate-job-instances
else
  $SCRIPTS/destroy-environment $ENV
fi
jenv=$JUJU_HOME/environments/$ENV.jenv
if [ -e $jenv ]; then rm $jenv; fi
