/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

// This file is maintained in the GoogleCloudPlatform/magic-modules repository and copied into the downstream provider repositories. Any changes to this file in the downstream will be overwritten.

package projects.reused

import MMUpstreamProjectId
import ProviderNameBeta
import ProviderNameGa
import ServiceSweeperCronName
import ServiceSweeperManualName
import SharedResourceNameVcr
import builds.*
import generated.PackagesListBeta
import generated.PackagesListGa
import generated.ServicesListBeta
import generated.ServicesListGa
import generated.SweepersListBeta
import generated.SweepersListGa
import jetbrains.buildServer.configs.kotlin.BuildType
import jetbrains.buildServer.configs.kotlin.Project
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot
import replaceCharsId

// MMUpstreamProjectInputs extends ReusableProjectInputs by adding the ability to
// pass in a second vcs root.
class MMUpstreamProjectInputs// Nullable inputs
    (
    override val parentProject: String,
    override val providerName: String,
    override val vcsRoot: GitVcsRoot,
    override val config: AccTestConfiguration,
    override val cron: NightlyTriggerConfiguration,
    override val projectName: String?,
    // This root is used to make the scheduled sweeping of the VCR project use the
    // downstream repo's code instead of the modular-magician fork.
    val cronSweeperVcsRoot: GitVcsRoot

    ) : ReusableProjectInputs(parentProject, providerName, vcsRoot, config, cron, projectName) {
}

fun mmUpstream(input: MMUpstreamProjectInputs): Project {

//    parentProject: String, providerName: String, vcsRoot: GitVcsRoot, cronSweeperVcsRoot: GitVcsRoot, config: AccTestConfiguration

    // Create unique ID for the dynamically-created project
    var projectId = "${input.parentProject}_${MMUpstreamProjectId}"
    projectId = replaceCharsId(projectId)

    // Shared resource allows ad hoc builds and sweeper builds to not clash
    var sharedResources: List<String> = listOf(SharedResourceNameVcr)

    // Create build configs for each package defined in packages.kt and services_ga.kt/services_beta.kt files
    val allPackages = getAllPackageInProviderVersion(input.providerName)
    val packageBuildConfigs = BuildConfigurationsForPackages(allPackages, input.providerName, projectId, input.vcsRoot, sharedResources, input.config)

    // Create build config for sweeping the VCR test project - everything except projects
    var sweepersList: Map<String,Map<String,String>>
    when(input.providerName) {
        ProviderNameGa -> sweepersList = SweepersListGa
        ProviderNameBeta -> sweepersList = SweepersListBeta
        else -> throw Exception("Provider name not supplied when generating a nightly test subproject")
    }
    val serviceSweeperManualConfig = BuildConfigurationForServiceSweeper(input.providerName, ServiceSweeperManualName, sweepersList, projectId, input.vcsRoot, sharedResources, input.config)

    val serviceSweeperCronConfig = BuildConfigurationForServiceSweeper(input.providerName, ServiceSweeperCronName, sweepersList, projectId, input.cronSweeperVcsRoot, sharedResources, input.config)
    val trigger  = input.cron
    serviceSweeperCronConfig.addTrigger(trigger) // Only the sweeper is on a schedule in this project

    return Project {
        id(projectId)
        name = (if (input.projectName != null) input.projectName!! else "Upstream MM Testing")
        description = "A project connected to the modular-magician/terraform-provider-${input.providerName} repository, to let users trigger ad-hoc builds against branches for PRs"

        // Register build configs in the project
        packageBuildConfigs.forEach { buildConfiguration: BuildType ->
            buildType(buildConfiguration)
        }
        buildType(serviceSweeperManualConfig)
        buildType(serviceSweeperCronConfig)

        params{
            configureGoogleSpecificTestParameters(input.config)
        }
    }
}

fun getAllPackageInProviderVersion(providerName: String): Map<String, Map<String,String>> {
    var allPackages: Map<String, Map<String, String>> = mapOf()
    if (providerName == ProviderNameGa){
        allPackages = PackagesListGa + ServicesListGa
    }
    if (providerName == ProviderNameBeta){
        allPackages = PackagesListBeta + ServicesListBeta
    }
    return allPackages
}