<script>
	import { onMount } from "svelte";
	import { stores } from "@sapper/app";
	import MainView from "./MainView/index.svelte";
	import Tags from "./Tags.svelte";
	import * as api from "api.js";

	export let p = 1;
	const { session } = stores();

	export function preload({ params }, { user }) {
		if (!user) {
			this.redirect(302, `/login`);
		}
		return;
	}
	let tab;

	onMount(async () => {
		({ proposal } = await api.get("proposal"));
	});
</script>

<svelte:head>
	<title>CreditTodxs</title>
</svelte:head>

<div class="home-page">
	<div
		class="bg-gradient-to-r from-indigo-400 to-blue-700 py-12 items-center text-center"
	>
		<h1 class="tracking-tight text-white text-2xl sm:text-4xl">
			CreditTodxs Ofertas de Crédito Para Você
		</h1>
	</div>
	<div class="container mx-auto">
		<div class="">
			<MainView {p} bind:tab />
		</div>
	</div>
</div>
