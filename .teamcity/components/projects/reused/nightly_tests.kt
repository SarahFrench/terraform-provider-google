/*
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

// This file is maintained in the GoogleCloudPlatform/magic-modules repository and copied into the downstream provider repositories. Any changes to this file in the downstream will be overwritten.

package projects.reused

import NightlyTestsProjectId
import ProviderNameBeta
import ProviderNameGa
import ServiceSweeperName
import SharedResourceNameBeta
import SharedResourceNameGa
import builds.*
import generated.SweepersListBeta
import generated.SweepersListGa
import jetbrains.buildServer.configs.kotlin.Project
import jetbrains.buildServer.configs.kotlin.vcs.GitVcsRoot
import replaceCharsId

open class ReusableProjectInputs(
    open val parentProject: String,
    open val providerName: String,
    open val vcsRoot: GitVcsRoot,
    open val config: AccTestConfiguration,
    open val cron: NightlyTriggerConfiguration,
    open val projectName: String
)

fun nightlyTests(input: ReusableProjectInputs): Project {

    // Create unique ID for the dynamically-created project
    var projectId = "${input.parentProject}_${NightlyTestsProjectId}"
    projectId = replaceCharsId(projectId)

    // Nightly test projects run all acceptance tests overnight
    // Here we ensure the project uses the appropriate Shared Resource to ensure no clashes between builds and/or sweepers
    var sharedResources: ArrayList<String>
    when(input.providerName) {
        ProviderNameGa -> sharedResources = arrayListOf(SharedResourceNameGa)
        ProviderNameBeta -> sharedResources = arrayListOf(SharedResourceNameBeta)
        else -> throw Exception("Provider name not supplied when generating a nightly test subproject")
    }

    // Create build configs to run acceptance tests for each package defined in packages.kt and services.kt files
    val allPackages = getAllPackageInProviderVersion(input.providerName)
    val packageBuildConfigs = BuildConfigurationsForPackages(allPackages, input.providerName, projectId, input.vcsRoot, sharedResources, input.config)

    // Add cron trigger to build configs for service packages
    packageBuildConfigs.forEach { buildConfiguration ->
        buildConfiguration.addTrigger(input.cron)
    }

    // Create build config for sweeping the nightly test project
    var sweepersList: Map<String,Map<String,String>>
    when(input.providerName) {
        ProviderNameGa -> sweepersList = SweepersListGa
        ProviderNameBeta -> sweepersList = SweepersListBeta
        else -> throw Exception("Provider name not supplied when generating a nightly test subproject")
    }
    val serviceSweeperConfig = BuildConfigurationForServiceSweeper(input.providerName, ServiceSweeperName, sweepersList, projectId, input.vcsRoot, sharedResources, input.config)

    // Add cron trigger to build config for service sweeper
    // Trigger must be scheduled after the service package builds
    val sweeperTrigger  = input.cron.clone()
    sweeperTrigger.startHour += 5
    serviceSweeperConfig.addTrigger(sweeperTrigger)

    return Project {
        id(projectId)
        name = input.projectName // Typically "Nightly Tests", but may take other values to ensure unique names within a project
        description = "A project connected to the hashicorp/terraform-provider-${input.providerName} repository, where scheduled nightly tests run and users can trigger ad-hoc builds"

        // Register build configs in the project
        packageBuildConfigs.forEach { buildConfiguration ->
            buildType(buildConfiguration)
        }
        buildType(serviceSweeperConfig)

        params{
            configureGoogleSpecificTestParameters(input.config)
        }
    }
}