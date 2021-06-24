<script>
	import * as api from "api.js";
	import { createEventDispatcher } from "svelte";
	export let offer;
	export let user;
	const dispatch = createEventDispatcher();

	async function toggleFavorite() {
		// optimistic UI
		if (offer.favorited) {
			offer.favoritesCount -= 1;
			offer.favorited = false;
		} else {
			offer.favoritesCount += 1;
			offer.favorited = true;
		}

		({ offer } = await (offer.favorited
			? api.post(
					`offers/${offer.slug}/favorite`,
					null,
					user && user.token
			  )
			: api.del(`offers/${offer.slug}/favorite`, user && user.token)));
	}
</script>

<div class="px-4 my-3 md:my-5">
	<div
		class="rounded shadow border-t-8 rounded border-blue-700 border break-words flex flex-col md:flex-row justify-between items-center"
	>
		<div>
			<a href="/offer/{offer.slug}" rel="prefetch" class="">
				<div class="bg-white dark:bg-gray-800 px-6 py-8">
					<h3
						class="text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl uppercase"
					>
						{offer.title}
					</h3>
					<p class="text-lg leading-6 font-medium text-gray-900">
						{offer.body}
					</p>
					<span>Ler mais...</span>
				</div>
			</a>
			<div
				class="text-gray-900 cursor-default hover:bg-indigo-500 hover:text-white select-none relative py-2 pl-3 pr-9 ml-3"
			>
				<a class="" href="/profile">
					<div class="flex items-center">
						<i
							class="ion-person flex-shrink-0 h-6 w-6 rounded-full"
						/>
						<span class="ml-3 block font-normal truncate"
							>&nbsp;{offer.author.username}</span
						>
						<span class="mx-1 text-sm italic"
							>{new Date(offer.createdAt).toDateString()}</span
						>
						<span class="inset-y-0 right-0 flex items-center pr-4"
							><svg
								class="h-5 w-5"
								xmlns="http://www.w3.org/2000/svg"
								viewBox="0 0 20 20"
								fill="currentColor"
								aria-hidden="true"
								><path
									fill-rule="evenodd"
									d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
									clip-rule="evenodd"
								/></svg
							>
						</span>
					</div>
				</a>
			</div>
		</div>
		<div class="p-3 text-center bg-gray-50 object-center">
			<button
				class="rounded-sm p-3 {offer.favorited
					? 'bg-blue-700 hover:bg-pink-500'
					: 'bg-gray-100 shadow hover:bg-black hover:text-white'}"
				on:click={toggleFavorite}
			>
				<i class="ion-heart" />&nbsp;{offer.favoritesCount}
			</button>
		</div>
	</div>
</div>
