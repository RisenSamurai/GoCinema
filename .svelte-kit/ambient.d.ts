
// this file is generated — do not edit it


/// <reference types="@sveltejs/kit" />

/**
 * Environment variables [loaded by Vite](https://vitejs.dev/guide/env-and-mode.html#env-files) from `.env` files and `process.env`. Like [`$env/dynamic/private`](https://kit.svelte.dev/docs/modules#$env-dynamic-private), this module cannot be imported into client-side code. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://kit.svelte.dev/docs/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://kit.svelte.dev/docs/configuration#env) (if configured).
 * 
 * _Unlike_ [`$env/dynamic/private`](https://kit.svelte.dev/docs/modules#$env-dynamic-private), the values exported from this module are statically injected into your bundle at build time, enabling optimisations like dead code elimination.
 * 
 * ```ts
 * import { API_KEY } from '$env/static/private';
 * ```
 * 
 * Note that all environment variables referenced in your code should be declared (for example in an `.env` file), even if they don't have a value until the app is deployed:
 * 
 * ```
 * MY_FEATURE_FLAG=""
 * ```
 * 
 * You can override `.env` values from the command line like so:
 * 
 * ```bash
 * MY_FEATURE_FLAG="enabled" npm run dev
 * ```
 */
declare module '$env/static/private' {
	export const ALLUSERSPROFILE: string;
	export const AMDRMPATH: string;
	export const APPCODE_VM_OPTIONS: string;
	export const APPDATA: string;
	export const ChocolateyInstall: string;
	export const ChocolateyLastPathUpdate: string;
	export const CLion: string;
	export const CLION_VM_OPTIONS: string;
	export const COLOR: string;
	export const CommonProgramFiles: string;
	export const CommonProgramW6432: string;
	export const COMPUTERNAME: string;
	export const ComSpec: string;
	export const CUDA_PATH: string;
	export const CUDA_PATH_V12_4: string;
	export const DATAGRIP_VM_OPTIONS: string;
	export const DATASPELL_VM_OPTIONS: string;
	export const DEVECOSTUDIO_VM_OPTIONS: string;
	export const DokanLibrary1: string;
	export const DokanLibrary1_LibraryPath_x64: string;
	export const DokanLibrary1_LibraryPath_x86: string;
	export const DriverData: string;
	export const EDITOR: string;
	export const EFC_8712: string;
	export const FIG_TERM: string;
	export const GATEWAY_VM_OPTIONS: string;
	export const GO111MODULE: string;
	export const GoLand: string;
	export const GOLAND_VM_OPTIONS: string;
	export const GOPATH: string;
	export const GOROOT: string;
	export const HOME: string;
	export const HOMEDRIVE: string;
	export const HOMEPATH: string;
	export const IDEA_VM_OPTIONS: string;
	export const INIT_CWD: string;
	export const INTELLIJ_TERMINAL_COMMAND_BLOCKS: string;
	export const JETBRAINSCLIENT_VM_OPTIONS: string;
	export const JETBRAINS_CLIENT_VM_OPTIONS: string;
	export const JETBRAINS_INTELLIJ_COMMAND_END_MARKER: string;
	export const LOCALAPPDATA: string;
	export const LOGONSERVER: string;
	export const MSMPI_BENCHMARKS: string;
	export const MSMPI_BIN: string;
	export const NODE: string;
	export const NODE_ENV: string;
	export const NODE_EXE: string;
	export const NPM_CLI_JS: string;
	export const npm_command: string;
	export const npm_config_cache: string;
	export const npm_config_globalconfig: string;
	export const npm_config_global_prefix: string;
	export const npm_config_init_module: string;
	export const npm_config_local_prefix: string;
	export const npm_config_node_gyp: string;
	export const npm_config_noproxy: string;
	export const npm_config_npm_version: string;
	export const npm_config_prefix: string;
	export const npm_config_userconfig: string;
	export const npm_config_user_agent: string;
	export const npm_execpath: string;
	export const npm_lifecycle_event: string;
	export const npm_lifecycle_script: string;
	export const npm_node_execpath: string;
	export const npm_package_json: string;
	export const npm_package_name: string;
	export const npm_package_version: string;
	export const NPM_PREFIX_NPM_CLI_JS: string;
	export const NUMBER_OF_PROCESSORS: string;
	export const OneDrive: string;
	export const OneDriveConsumer: string;
	export const OPENSSL_CONF: string;
	export const OS: string;
	export const Path: string;
	export const PATHEXT: string;
	export const PhpStorm: string;
	export const POWERSHELL_DISTRIBUTION_CHANNEL: string;
	export const PROCESSOR_ARCHITECTURE: string;
	export const PROCESSOR_IDENTIFIER: string;
	export const PROCESSOR_LEVEL: string;
	export const PROCESSOR_REVISION: string;
	export const PROCESS_LAUNCHED_BY_CW: string;
	export const PROCESS_LAUNCHED_BY_Q: string;
	export const ProgramData: string;
	export const ProgramFiles: string;
	export const ProgramW6432: string;
	export const PROMPT: string;
	export const PSExecutionPolicyPreference: string;
	export const PSModulePath: string;
	export const PUBLIC: string;
	export const PyCharm: string;
	export const PYCHARM_VM_OPTIONS: string;
	export const RIDER_VM_OPTIONS: string;
	export const RUBYMINE_VM_OPTIONS: string;
	export const SESSIONNAME: string;
	export const STUDIO_VM_OPTIONS: string;
	export const SystemDrive: string;
	export const SystemRoot: string;
	export const TEMP: string;
	export const TERMINAL_EMULATOR: string;
	export const TERM_SESSION_ID: string;
	export const TMP: string;
	export const USERDOMAIN: string;
	export const USERDOMAIN_ROAMINGPROFILE: string;
	export const USERNAME: string;
	export const USERPROFILE: string;
	export const WEBIDE_VM_OPTIONS: string;
	export const WebStorm: string;
	export const WEBSTORM_VM_OPTIONS: string;
	export const windir: string;
	export const __PSLockDownPolicy: string;
}

/**
 * Similar to [`$env/static/private`](https://kit.svelte.dev/docs/modules#$env-static-private), except that it only includes environment variables that begin with [`config.kit.env.publicPrefix`](https://kit.svelte.dev/docs/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Values are replaced statically at build time.
 * 
 * ```ts
 * import { PUBLIC_BASE_URL } from '$env/static/public';
 * ```
 */
declare module '$env/static/public' {
	
}

/**
 * This module provides access to runtime environment variables, as defined by the platform you're running on. For example if you're using [`adapter-node`](https://github.com/sveltejs/kit/tree/main/packages/adapter-node) (or running [`vite preview`](https://kit.svelte.dev/docs/cli)), this is equivalent to `process.env`. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://kit.svelte.dev/docs/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://kit.svelte.dev/docs/configuration#env) (if configured).
 * 
 * This module cannot be imported into client-side code.
 * 
 * Dynamic environment variables cannot be used during prerendering.
 * 
 * ```ts
 * import { env } from '$env/dynamic/private';
 * console.log(env.DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 * 
 * > In `dev`, `$env/dynamic` always includes environment variables from `.env`. In `prod`, this behavior will depend on your adapter.
 */
declare module '$env/dynamic/private' {
	export const env: {
		ALLUSERSPROFILE: string;
		AMDRMPATH: string;
		APPCODE_VM_OPTIONS: string;
		APPDATA: string;
		ChocolateyInstall: string;
		ChocolateyLastPathUpdate: string;
		CLion: string;
		CLION_VM_OPTIONS: string;
		COLOR: string;
		CommonProgramFiles: string;
		CommonProgramW6432: string;
		COMPUTERNAME: string;
		ComSpec: string;
		CUDA_PATH: string;
		CUDA_PATH_V12_4: string;
		DATAGRIP_VM_OPTIONS: string;
		DATASPELL_VM_OPTIONS: string;
		DEVECOSTUDIO_VM_OPTIONS: string;
		DokanLibrary1: string;
		DokanLibrary1_LibraryPath_x64: string;
		DokanLibrary1_LibraryPath_x86: string;
		DriverData: string;
		EDITOR: string;
		EFC_8712: string;
		FIG_TERM: string;
		GATEWAY_VM_OPTIONS: string;
		GO111MODULE: string;
		GoLand: string;
		GOLAND_VM_OPTIONS: string;
		GOPATH: string;
		GOROOT: string;
		HOME: string;
		HOMEDRIVE: string;
		HOMEPATH: string;
		IDEA_VM_OPTIONS: string;
		INIT_CWD: string;
		INTELLIJ_TERMINAL_COMMAND_BLOCKS: string;
		JETBRAINSCLIENT_VM_OPTIONS: string;
		JETBRAINS_CLIENT_VM_OPTIONS: string;
		JETBRAINS_INTELLIJ_COMMAND_END_MARKER: string;
		LOCALAPPDATA: string;
		LOGONSERVER: string;
		MSMPI_BENCHMARKS: string;
		MSMPI_BIN: string;
		NODE: string;
		NODE_ENV: string;
		NODE_EXE: string;
		NPM_CLI_JS: string;
		npm_command: string;
		npm_config_cache: string;
		npm_config_globalconfig: string;
		npm_config_global_prefix: string;
		npm_config_init_module: string;
		npm_config_local_prefix: string;
		npm_config_node_gyp: string;
		npm_config_noproxy: string;
		npm_config_npm_version: string;
		npm_config_prefix: string;
		npm_config_userconfig: string;
		npm_config_user_agent: string;
		npm_execpath: string;
		npm_lifecycle_event: string;
		npm_lifecycle_script: string;
		npm_node_execpath: string;
		npm_package_json: string;
		npm_package_name: string;
		npm_package_version: string;
		NPM_PREFIX_NPM_CLI_JS: string;
		NUMBER_OF_PROCESSORS: string;
		OneDrive: string;
		OneDriveConsumer: string;
		OPENSSL_CONF: string;
		OS: string;
		Path: string;
		PATHEXT: string;
		PhpStorm: string;
		POWERSHELL_DISTRIBUTION_CHANNEL: string;
		PROCESSOR_ARCHITECTURE: string;
		PROCESSOR_IDENTIFIER: string;
		PROCESSOR_LEVEL: string;
		PROCESSOR_REVISION: string;
		PROCESS_LAUNCHED_BY_CW: string;
		PROCESS_LAUNCHED_BY_Q: string;
		ProgramData: string;
		ProgramFiles: string;
		ProgramW6432: string;
		PROMPT: string;
		PSExecutionPolicyPreference: string;
		PSModulePath: string;
		PUBLIC: string;
		PyCharm: string;
		PYCHARM_VM_OPTIONS: string;
		RIDER_VM_OPTIONS: string;
		RUBYMINE_VM_OPTIONS: string;
		SESSIONNAME: string;
		STUDIO_VM_OPTIONS: string;
		SystemDrive: string;
		SystemRoot: string;
		TEMP: string;
		TERMINAL_EMULATOR: string;
		TERM_SESSION_ID: string;
		TMP: string;
		USERDOMAIN: string;
		USERDOMAIN_ROAMINGPROFILE: string;
		USERNAME: string;
		USERPROFILE: string;
		WEBIDE_VM_OPTIONS: string;
		WebStorm: string;
		WEBSTORM_VM_OPTIONS: string;
		windir: string;
		__PSLockDownPolicy: string;
		[key: `PUBLIC_${string}`]: undefined;
		[key: `${string}`]: string | undefined;
	}
}

/**
 * Similar to [`$env/dynamic/private`](https://kit.svelte.dev/docs/modules#$env-dynamic-private), but only includes variables that begin with [`config.kit.env.publicPrefix`](https://kit.svelte.dev/docs/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Note that public dynamic environment variables must all be sent from the server to the client, causing larger network requests — when possible, use `$env/static/public` instead.
 * 
 * Dynamic environment variables cannot be used during prerendering.
 * 
 * ```ts
 * import { env } from '$env/dynamic/public';
 * console.log(env.PUBLIC_DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 */
declare module '$env/dynamic/public' {
	export const env: {
		[key: `PUBLIC_${string}`]: string | undefined;
	}
}
