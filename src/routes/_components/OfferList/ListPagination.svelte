<script>
	export let offersCount;
	export let page;

	let range;

	$: {
		range = [];
		for (let i = 0; i < Math.ceil(offersCount / 10); ++i) {
			range.push(i);
		}
	}
</script>

{#if offersCount > 10}
	<div class="container mx-auto">
		<div
			class="bg-white px-4 py-3 flex items-center justify-between border-t border-gray-200 sm:px-6"
		>
			<div
				class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between"
			>
				<div>
					<p class="text-sm text-gray-700">
						Mostrando
						<span class="font-medium">{offersCount}</span>
						resultados
					</p>
				</div>
				<div>
					<nav
						class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px"
						aria-label="Pagination"
					>
						{#each range as v}
							<a
								class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50 {v ===
								page
									? 'bg-gray-400'
									: ''}"
								href="/{v ? v + 1 : ''}"
								>{v + 1}
							</a>
						{/each}
					</nav>
				</div>
			</div>
		</div>
	</div>
{/if}
