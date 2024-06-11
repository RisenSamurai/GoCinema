<script>

import HomeMoviesSlider from "$lib/components/HomeMoviesSlider.svelte";
</script>

<svelte:head>
    <title>GoCinema | Home</title>
    <meta name="description" content="An online cinema where you can watch the best movies and series of the world!">
</svelte:head>


<main class="flex flex-col px-2 bg-cinema-primary">
    <HomeMoviesSlider header="New Movies"/>
    <HomeMoviesSlider header="Hot News" background="bg-white"/>
    <HomeMoviesSlider header="Popular Series"/>
</main>


