import { readable } from 'svelte/store'


export const curPath = readable("/", function pathname() {
    return window.location.pathname;
});

