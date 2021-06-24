<script>
	export let error, status;

	let dev = process.env.NODE_ENV === "development";
</script>

<svelte:head>
	<title>{status}</title>
</svelte:head>

<div class="bg-indigo-900 relative overflow-hidden h-screen">
	<img
		src="/img/8.svg"
		class="absolute h-full w-full object-cover"
		alt="error background"
	/>
	<div class="inset-0 bg-black opacity-25 absolute" />
	<div
		class="container mx-auto px-6 md:px-12 relative z-10 flex items-center py-32 xl:py-40"
	>
		<div class="w-full font-mono flex flex-col items-center relative z-10">
			<h1
				class="font-extrabold text-5xl text-center text-white leading-tight mt-4"
			>
				{status}
			</h1>
			<h3 class="text-4xl text-center text-white leading-tight mt-4">
				{error.message}
			</h3>
			<p class="font-extrabold text-1xl my-44 text-white animate-bounce">
				{#if dev && error.stack}
					<pre>{error.stack}</pre>
				{/if}
			</p>
		</div>
	</div>
</div>
