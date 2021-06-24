<script context="module">
	import * as api from "api.js";

	export async function preload({ params }) {
		const { slug } = params;
		const { offer } = await api.get(`offers/${params.slug}`, null);

		return { offer, slug };
	}
</script>

<script>
	import { onMount } from "svelte";
	import { stores } from "@sapper/app";
	import marked from "marked";

	export let offer;

	const { session } = stores();
	let user = $session.user;

	$: markup = marked(offer.body);

	$: canModify = user && offer.author.username === user.username;

	function remove() {
		api.del(`offers/${offer.slug}`, user && user.token);
		goto("/");
	}
</script>

<svelte:head>
	<title>{offer.title}</title>
</svelte:head>

<div>
	<div class="block justify-center">
		<div
			class="lg:flex lg:items-center lg:justify-between border rounded py-3 bg-gray-300 py-3 md:py-5 px-4 sm:px-6 lg:px-8"
		>
			<div class="flex-1 min-w-0">
				<h2
					class="text-2xl font-bold leading-7 text-gray-900 sm:text-3xl sm:truncate"
				>
					{offer.title}
				</h2>
				<div
					class="mt-1 flex flex-col sm:flex-row sm:flex-wrap sm:mt-0 sm:space-x-6"
				>
					<a
						href="/profile/@{user.name}"
						class="mt-2 flex items-center text-sm text-gray-500"
					>
						<svg
							class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
							xmlns="http://www.w3.org/2000/svg"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M6 6V5a3 3 0 013-3h2a3 3 0 013 3v1h2a2 2 0 012 2v3.57A22.952 22.952 0 0110 13a22.95 22.95 0 01-8-1.43V8a2 2 0 012-2h2zm2-1a1 1 0 011-1h2a1 1 0 011 1v1H8V5zm1 5a1 1 0 011-1h.01a1 1 0 110 2H10a1 1 0 01-1-1z"
								clip-rule="evenodd"
							/>
							<path
								d="M2 13.692V16a2 2 0 002 2h12a2 2 0 002-2v-2.308A24.974 24.974 0 0110 15c-2.796 0-5.487-.46-8-1.308z"
							/>
						</svg>
						{user.name}
					</a>
					<div class="mt-2 flex items-center text-sm text-gray-500">
						<svg
							class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"
							xmlns="http://www.w3.org/2000/svg"
							viewBox="0 0 20 20"
							fill="currentColor"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z"
								clip-rule="evenodd"
							/>
						</svg>
						{new Date(offer.createdAt).toDateString()}
					</div>
				</div>
			</div>
			{#if canModify}
				<div class="mt-5 flex lg:mt-0 lg:ml-4">
					<span class="">
						<a
							href="/user/{offer.slug}"
							type="button"
							class="inline-flex items-center px-4 py-2 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
						>
							<i class="ion-edit" />Aceitar
						</a>
					</span>
				</div>
			{/if}
		</div>
		<div class="py-3 md:py-5 pb-5 md:pb-9 px-4 sm:px-6 lg:px-8">
			<p class="text-2xl text-gray-700 md:text-3xl">{@html markup}</p>
		</div>
		<hr />
	</div>
</div>
