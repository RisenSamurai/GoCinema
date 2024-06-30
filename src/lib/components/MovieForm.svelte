<script>

    import Button from "$lib/components/Button.svelte";

    let message = "";
    let director = "";
    let directors = [];
    let writers = [];
    let producers = [];
    let editors = [];
    let cameras = [];
    let genres = [];


    function pushDirector(director) {
        directors = [...directors, director]
    }

    function removeDirector(index) {
        directors = directors.filter((_, i) => i !== index);
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
        <input bind:value={director} class="p-1 rounded-lg mb-2" type="text" name="director" id="director" required>

        {#if directors.length > 0}
            <div class="flex flex-col">

                {#each directors as director, index}

                    <span class="text-cinema-text">{index}{director}</span>
                    <button type="button" class="text-cinema-text font-bold text-2xl"
                            on:click={() => removeDirector(index)}>X</button>
                    {/each}

            </div>
        {/if}
        <Button name="Push" on:click={() => pushDirector(director)} />

    </fieldset>

</form>