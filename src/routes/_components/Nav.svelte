<script>
	import { stores } from "@sapper/app";

	const { page, session } = stores();
	function logout() {
		$session.user = null;
	}

	let show = false;
	function dropdown() {
		if (show) {
			document.getElementById("mobile-menu").classList.value = "hidden";
		} else {
			document.getElementById("mobile-menu").classList.value = "";
		}
		show = !show;
	}
</script>

{#if $session.user}
	<div>
		<nav class="bg-gray-800">
			<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div class="flex items-center justify-between h-16">
					<div class="flex items-center">
						<div class="flex-shrink-0">
							<a href="/"
								><img
									class="h-8"
									src="logo.png"
									alt="CreditTodxs"
								/></a
							>
						</div>
						<div class="hidden md:block">
							<div class="ml-10 flex items-baseline space-x-4">
								<!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
								<a
									href="/"
									rel="prefetch"
									class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
									>Ofertas</a
								>
							</div>
						</div>
					</div>
					<div class="hidden md:block">
						<div class="ml-4 flex items-center md:ml-6">
							<!-- Profile dropdown -->
							<div class="ml-3 relative dropdown">
								<button
									type="button"
									class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
									aria-expanded="false"
									aria-haspopup="true"
								>
									<span class="sr-only"
										>Abrir Menu de Usuário</span
									>
									<div class="">
										<i class="ion-android-person" />
									</div>
									<div>{$session.user.username}</div>
								</button>
								<div
									class="dropdown-menu origin-top-right absolute right-0 w-48 rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 focus:outline-none hidden"
									role="menu"
									aria-orientation="vertical"
									aria-labelledby="user-menu"
								>
									<button
										class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
										role="menuitem"
										on:click={logout}>Sair</button
									>
								</div>
							</div>
						</div>
					</div>

					<div class="-mr-2 flex md:hidden">
						<!-- Mobile menu button -->
						<button
							type="button"
							class="bg-gray-800 inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
							aria-controls="mobile-menu"
							aria-expanded="false"
							on:click={dropdown}
						>
							<span class="sr-only">Abrir Menu</span>
							<!--
              Heroicon name: outline/menu

              Menu open: "hidden", Menu closed: "block"
            -->
							<svg
								class="block h-6 w-6"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M4 6h16M4 12h16M4 18h16"
								/>
							</svg>
							<!--
              Heroicon name: outline/x

              Menu open: "block", Menu closed: "hidden"
            -->
							<svg
								class="hidden h-6 w-6"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
								aria-hidden="true"
							>
								<path
									stroke-linecap="round"
									stroke-linejoin="round"
									stroke-width="2"
									d="M6 18L18 6M6 6l12 12"
								/>
							</svg>
						</button>
					</div>
				</div>
			</div>

			<!-- Mobile menu, show/hide based on menu state. -->
			<div class="hidden" id="mobile-menu">
				<div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
					<!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
					<a
						href="/"
						rel="prefetch"
						class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
						>Ofertas</a
					>
				</div>
				<div class="pt-4 pb-3 border-t border-gray-700">
					<div class="mt-3 px-2 space-y-1">
						<button
							class="block px-3 py-2 rounded-md text-base font-medium text-gray-400 hover:text-white hover:bg-gray-700"
							on:click={logout}>Sair<button /></button
						>
					</div>
				</div>
			</div>
		</nav>
	</div>
{:else}
	<div>
		<nav class="bg-gray-800">
			<div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div class="flex items-center justify-between h-16">
					<div class="flex items-center">
						<div class="flex-shrink-0">
							<a href="/"
								><img
									class="h-8"
									src="logo.png"
									alt="CreditTodxs"
								/></a
							>
						</div>
						<div class="hidden md:block">
							<div class="ml-10 flex items-baseline space-x-4">
								<!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
								<a
									rel="prefetch"
									href="/login"
									class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
									class:active={$page.path === "/login"}
								>
									Logar
								</a>

								<a
									rel="prefetch"
									href="/register"
									class="text-gray-300 hover:bg-gray-700 hover:text-white px-3 py-2 rounded-md text-sm font-medium"
									class:active={$page.path === "/register"}
								>
									<i class="ion-android-person-add" />
									Registrar
								</a>
							</div>
						</div>
					</div>

					<!-- Mobile menu, show/hide based on menu state. -->
					<div class="md:hidden" id="mobile-menu">
						<div class="px-2 pt-2 pb-3 space-y-1 sm:px-3">
							<!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->
							<a
								href="/login"
								rel="prefetch"
								class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
								class:active={$page.path === "/login"}>Logar</a
							>

							<a
								href="/register"
								rel="prefetch"
								class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium"
								class:active={$page.path === "/register"}
								><i class="ion-android-person-add" />
								Registrar</a
							>
						</div>
					</div>
				</div>
			</div>
		</nav>
	</div>
{/if}
