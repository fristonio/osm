package main

import "strings"

// validateCLIParams contains all checks necessary that various permutations of the CLI flags are consistent
// Side effects: uses log.Fatal() resulting in os.Exit(255)
func validateCLIParams() {
	if _, ok := certManagers[certificateManagerKind(*certManagerKind)]; !ok {
		log.Fatal().Msgf("Certificate manager %s is not one of possible options: %s", *certManagerKind, strings.Join(getPossibleCertManagers(), ", "))
	}

	if *certManagerKind == vaultKind {
		if *vaultToken == "" {
			log.Fatal().Msg("Empty Hashi Vault token.")
		}
	}

	if meshName == "" {
		log.Fatal().Msg("Please specify the mesh name using --mesh-name")
	}

	if osmNamespace == "" {
		log.Fatal().Msg("Please specify the OSM namespace using --osm-namespace")
	}

	if injectorConfig.InitContainerImage == "" {
		log.Fatal().Msg("Please specify the init container image using --init-container-image")
	}

	if injectorConfig.SidecarImage == "" {
		log.Fatal().Msg("Please specify the sidecar image using --sidecar-image")
	}

	if webhookName == "" {
		log.Fatal().Msgf("Invalid --webhook-name value: '%s'", webhookName)
	}
}
