<script context="module">
	import Home from "./_components/Home.svelte";
	export function preload({ params }, { user }) {
		if (!user) {
			this.redirect(302, `/login`);
		}
		return;
	}
</script>

<Home p={1} />
