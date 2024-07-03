/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

// This file is controlled by MMv1, any changes made here will be overwritten

package projects.feature_branches

import ProviderNameBeta
import ProviderNameGa
import builds.*
import generated.PackagesListBeta
import generated.PackagesListGa
import jetbrains.buildServer.configs.kotlin.Project
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot
import projects.reused.ReusableProjectInputs
import projects.reused.nightlyTests
import replaceCharsId
import vcs_roots.HashiCorpVCSRootGa
import vcs_roots.ModularMagicianVCSRootBeta
import vcs_roots.ModularMagicianVCSRootGa

const val branchName = "FEATURE-BRANCH-major-release-6.0.0"


// VCS Roots specifically for pulling code from the feature branches in the downstream and upstream repos
object HashicorpVCSRootGa_featureBranchMajorRelease600: GitVcsRoot({
    name = "VCS root for the hashicorp/terraform-provider-${ProviderNameGa} repo @ refs/heads/${branchName}"
    url = "https://github.com/hashicorp/terraform-provider-${ProviderNameGa}"
    branch = "refs/heads/${branchName}"
    branchSpec = "" // empty as we'll access no other branches
})

object HashicorpVCSRootBeta_featureBranchMajorRelease600: GitVcsRoot({
    name = "VCS root for the hashicorp/terraform-provider-${ProviderNameBeta} repo @ refs/heads/${branchName}"
    url = "https://github.com/hashicorp/terraform-provider-${ProviderNameBeta}"
    branch = "refs/heads/${branchName}"
    branchSpec = "" // empty as we'll access no other branches
})
fun featureBranchMajorRelease600_Project(allConfig: AllContextParameters): Project {

    val projectId = replaceCharsId(branchName)
    val gaProjectId = replaceCharsId(branchName + "_GA")
    val betaProjectId= replaceCharsId(branchName + "_BETA")

    // Get config for using the GA and VCR identities
    val gaConfig = getGaAcceptanceTestConfig(allConfig)
    val betaConfig = getBetaAcceptanceTestConfig(allConfig)

    val gaProjectInputs = ReusableProjectInputs(
        parentProject = gaProjectId,
        providerName = ProviderNameGa,
        vcsRoot = HashicorpVCSRootGa_featureBranchMajorRelease600,
        config= gaConfig,
        cron= NightlyTriggerConfiguration(),
        projectName = "Nightly Tests (GA)"
    )
    val betaProjectInputs = ReusableProjectInputs(
        parentProject = betaProjectId,
        providerName = ProviderNameBeta,
        vcsRoot = HashicorpVCSRootBeta_featureBranchMajorRelease600,
        config= betaConfig,
        cron= NightlyTriggerConfiguration(),
        projectName = "Nightly Tests (Beta)"
    )

    return Project{
        id(projectId)
        name = "Testing for the 6.0.0 major release"
        description = "Subproject for testing feature branch $branchName"

        // Register feature branch-specific VCS roots in the project
        vcsRoot(HashicorpVCSRootGa_featureBranchMajorRelease600)
        vcsRoot(HashicorpVCSRootBeta_featureBranchMajorRelease600)

        // Nested Nightly Test project that uses hashicorp/terraform-provider-google
        subProject(nightlyTests(gaProjectInputs))

        // Nested Nightly Test project that uses hashicorp/terraform-provider-google-beta
        subProject(nightlyTests(betaProjectInputs))


        params {
            readOnlySettings()
        }
    }
}