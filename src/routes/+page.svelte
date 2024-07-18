<script>

import HomeMoviesSlider from "$lib/components/HomeMoviesSlider.svelte";
export let data;
</script>

<svelte:head>
    <title>GoCinema | Home</title>
    <meta name="description" content="An online cinema where you can watch the best movies and series of the world!">
</svelte:head>


<main class="flex flex-col px-2 bg-cinema-primary">
    <HomeMoviesSlider header="New Movies" type="movies" items={data.items}/>
    <HomeMoviesSlider header="Hot News" type="news" />
    <HomeMoviesSlider header="TV Shows" type="series"/>
</main>


