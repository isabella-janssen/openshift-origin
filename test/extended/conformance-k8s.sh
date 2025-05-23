#!/bin/bash
#
# Runs the Kubernetes conformance suite against an OpenShift cluster
#
# Test prerequisites:
#
# * all nodes that users can run workloads under marked as schedulable
#
source "$(dirname "${BASH_SOURCE}")/../../hack/lib/init.sh"

# Check inputs
if [[ -z "${KUBECONFIG-}" ]]; then
  os::log::fatal "KUBECONFIG must be set to a root account"
fi
test_report_dir="${ARTIFACT_DIR}"
mkdir -p "${test_report_dir}"

cat <<END > "${test_report_dir}/README.md"
This conformance report is generated by the OpenShift CI infrastructure. The canonical source location for this test script is located at https://github.com/openshift/origin/blob/main/test/extended/conformance-k8s.sh

This file was generated by:

  Commit $( git rev-parse HEAD || "<commit>" )
  Tag    $( git describe || "<tag>" )

To recreate these results

1. Install an [OpenShift cluster](https://docs.openshift.com/container-platform/latest/install_config/install/advanced_install.html)
2. Retrieve a \`.kubeconfig\` file with administrator credentials on that cluster and set the environment variable KUBECONFIG

    export KUBECONFIG=PATH_TO_KUBECONFIG

3. Clone the OpenShift source repository and change to that directory:

    git clone https://github.com/openshift/origin.git
    cd origin

4. Place the \`oc\` binary for that cluster in your PATH
5. Run the conformance test:

    test/extended/conformance-k8s.sh

Nightly conformance tests are run against release branches and reported https://openshift-gce-devel.appspot.com/builds/origin-ci-test/logs/periodic-ci-origin-conformance-k8s/
END

version="${KUBERNETES_VERSION:-release-1.18}"
kubernetes="${KUBERNETES_ROOT:-${OS_ROOT}/../../../k8s.io/kubernetes}"
if [[ -d "${kubernetes}" ]]; then
  git fetch origin --tags
else
  if [[ -n "${KUBERNETES_ROOT-}" ]]; then
    os::log::fatal "Cannot find Kubernetes source directory, set KUBERNETES_ROOT"
  fi
  kubernetes="${OS_ROOT}/_output/components/kubernetes"
  if [[ ! -d "${kubernetes}" ]]; then
    mkdir -p "$( dirname "${kubernetes}" )"
    os::log::info "Cloning Kubernetes source"
    git clone "https://github.com/kubernetes/kubernetes.git" -b "${version}" "${kubernetes}" # --depth=1 unfortunately we need history info as well
  fi
fi

os::log::info "Running Kubernetes conformance suite for ${version}"

# Execute OpenShift prerequisites
# Disable container security
oc adm policy add-scc-to-group privileged system:authenticated system:serviceaccounts
oc adm policy add-scc-to-group anyuid system:authenticated system:serviceaccounts
# Mark the master nodes as unschedulable so tests ignore them
oc get nodes -o name -l 'node-role.kubernetes.io/master' | xargs -L1 oc adm cordon
unschedulable="$( ( oc get nodes -o name -l 'node-role.kubernetes.io/master'; ) | wc -l )"
# TODO: undo these operations

# Execute Kubernetes prerequisites
pushd "${kubernetes}" > /dev/null
git checkout "${version}"
make WHAT=cmd/kubectl
make WHAT=test/e2e/e2e.test
make WHAT=vendor/github.com/onsi/ginkgo/v2/ginkgo
export PATH="${kubernetes}/_output/local/bin/$( os::build::host_platform ):${PATH}"

kubectl version  > "${test_report_dir}/version.txt"
echo "-----"    >> "${test_report_dir}/version.txt"
oc version      >> "${test_report_dir}/version.txt"

# Run the test, serial tests first, then parallel

rc=0

ginkgo \
  -nodes 1 -noColor '-focus=(\[Conformance\].*\[Serial\]|\[Serial\].*\[Conformance\])' $( which e2e.test ) -- \
  -report-dir "${test_report_dir}" \
  -allowed-not-ready-nodes ${unschedulable} \
  2>&1 | tee -a "${test_report_dir}/e2e.log" || rc=1

rename -v junit_ junit_serial_ ${test_report_dir}/junit*.xml

ginkgo \
  -nodes 4 -noColor '-skip=\[Serial\]' '-focus=\[Conformance\]' $( which e2e.test ) -- \
  -report-dir "${test_report_dir}" \
  -allowed-not-ready-nodes ${unschedulable} \
  2>&1 | tee -a "${test_report_dir}/e2e.log" || rc=1

echo
echo "Run complete, results in ${test_report_dir}"

exit $rc
