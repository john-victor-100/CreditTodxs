<script context="module">
	export function preload({ params }, { user }) {
		if (!user) {
			this.redirect(302, `/login`);
			return;
		}
		return {
			p: +params.p,
		};
	}
</script>

<script>
	import Home from "./_components/Home.svelte";
	export let p;
</script>

<Home {p} />
