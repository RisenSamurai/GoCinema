<script>

    import Button from "$lib/components/Button.svelte";

    let message = "";
    let director = "";
    let directors = [];




    function pushDirector(event) {

    }

</script>



<h2 class="text-2xl text-white">{#if message}{message}{/if}</h2>
<form on:submit class="flex flex-col bg-cinema-secondary w-full">
    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Basic Info</legend>
        <label class="text-cinema-text" for="movie-name">Movie Name</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="movie-name" id="movie-name">
        <label class="text-cinema-text" for="year">Year</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="lastName" id="year">
    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Film Crew</legend>
        <label class="text-cinema-text" for="director">Director</label>
        <input bind:value={director} class="p-1 rounded-lg mb-2" type="text" name="director" id="director">
        {#if director.length > 0}
            <div class="flex flex-col">

            </div>
        {/if}
        <Button name="Push" on:click={pushDirector} />

    </fieldset>

</form>