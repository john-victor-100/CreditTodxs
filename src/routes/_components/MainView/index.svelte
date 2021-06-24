<script>
	import { stores } from "@sapper/app";
	import OfferList from "../OfferList/index.svelte";

	export let p;

	const { session } = stores();
</script>

<OfferList {p} />
