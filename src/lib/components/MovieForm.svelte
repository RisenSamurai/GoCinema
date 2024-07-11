<script>

    import Button from "$lib/components/Button.svelte";

    let message = "";
    let director = "";
    let writer = "";
    let producer = "";
    let editor = "";
    let camera = "";
    let genre = "";
    let country = "";
    let actor = "";

    let directors = [];
    let writers = [];
    let producers = [];
    let editors = [];
    let cameras = [];
    let genres = [];
    let countries = [];
    let images = [];
    let actors = [];
    let previews = [];



    function pushDirector(person, type) {
        event.preventDefault();

        switch (type) {
            case "d": {
                directors = [...directors, person];
                director = "";
            } break;
            case "w": {
                writers = [...writers, person];
                writer = "";
            } break;
            case "p": {
                producers = [...producers, person];
                producer = "";
            } break;
            case "e": {
                editors = [...editors, person];
                editor = "";
            } break;
            case "c": {
                cameras = [...cameras, person];
                camera = "";
            } break;
            case "g": {
                genres = [...genres, person];
                genre = "";
            } break;
            case "a": {
                actors = [...actors, person];
                actor = "";
            } break;
            case "country": {
                countries = [...countries, person]
            } break;

            default: "Error"; break;
        }


    }

    function removeDirector(index, type) {

        event.preventDefault();

        switch (type) {
            case "d": {
                directors = directors.filter((_, i) => i !== index);
            } break;

            case "w": {
                writers = writers.filter((_, i) => i !== index);
            } break;

            case "p": {
                producers = producers.filter((_, i) => i !== index);
            } break;

            case "e": {
                editors = editors.filter((_, i) => i !== index);
            } break;

            case "c": {
                cameras = cameras.filter((_, i) => i !== index);
            } break;

            case "g": {
                genres = genres.filter((_, i) => i !== index);
            } break;

            case "a": {
                actors = actors.filter((_, i) => i !== index);
            } break;
            case "country": {
                countries = countries.filter((_, i) => i !== index);
            }break;
            default: "Error"; break;
        }


    }

    function pushImages(event) {

        const input = event.target;
        if (input && input.files.length > 0) {
            const newImages = Array.from(input.files);
            const newPreviews = newImages.map(file => URL.createObjectURL(file));
            images = [...images, ...newImages];
            previews = [...previews, ...newPreviews];
        }
    }

    function deleteImage(index) {
        images = images.filter((_, i) => i !== index);
        previews = previews.filter((_, i) => i !== index);
    }


   async function sendData(event) {

        event.preventDefault();

        const form = event.target;
        const formData = new FormData(form);

        images.forEach(image => {
            formData.append('images', image);
        });

       formData.append("directors", JSON.stringify(directors));
       formData.append("writers", JSON.stringify(writers));
       formData.append("producers", JSON.stringify(producers));
       formData.append("editors", JSON.stringify(editors));
       formData.append("cameras", JSON.stringify(cameras));
       formData.append("genres", JSON.stringify(genres));
       formData.append("countries", JSON.stringify(countries));
       formData.append("actors", JSON.stringify(actors));

        try {
            const response = await fetch("http://localhost:8000/add-movie", {
            method: 'POST',
            body: formData,
        });

        const data = await response.json();

        if (data.success) {
            console.log('Actor added successfully:', message = data.message);
        } else {
            console.error('Error adding actor:', message = data.message);
        }



        } catch(e) {
            console.error(e);
        }


    }


</script>



<h2 class="text-2xl text-white">{#if message}{message}{/if}</h2>
<form on:submit={sendData} class="flex flex-col bg-cinema-secondary w-full">
    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Basic Info</legend>
        <label class="text-cinema-text" for="movie-name">Movie Name</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="movie-name" id="movie-name">
        <label class="text-cinema-text" for="year">Year</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="year" id="year">
    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Film Crew</legend>
        <label class="text-cinema-text" for="director">Director</label>
        <input bind:value={director} class="p-1 rounded-lg mb-2" type="text" name="director" id="director">

        {#if directors.length > 0}
            <div class="flex flex-wrap">

                {#each directors as director, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{director}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeDirector(index, "d")}>X</button>
                    </div>
                    {/each}

            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushDirector(director, "d")} />

        <label class="text-cinema-text" for="director">Writers</label>
        <input bind:value={writer} class="p-1 rounded-lg mb-2" type="text" name="writers" id="director">

        {#if directors.length > 0}
            <div class="flex flex-wrap">

                {#each writers as writer, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{writer}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeDirector(index, "w")}>X</button>
                    </div>
                {/each}

            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushDirector(writer, "w")} />


        <label class="text-cinema-text" for="director">Producers</label>
        <input bind:value={producer} class="p-1 rounded-lg mb-2" type="text" name="producers" id="director">

        {#if producers.length > 0}
            <div class="flex flex-wrap">

                {#each producers as writer, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{writer}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeDirector(index, "p")}>X</button>
                    </div>
                {/each}

            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushDirector(producer, "p")} />

        <label class="text-cinema-text" for="director">Editors</label>
        <input bind:value={editor} class="p-1 rounded-lg mb-2" type="text" name="editors" id="director">

        {#if editors.length > 0}
            <div class="flex flex-wrap">

                {#each editors as editor, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{editor}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeDirector(index, "e")}>X</button>
                    </div>
                {/each}

            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushDirector(editor, "e")} />

        <label class="text-cinema-text" for="director">Cameras</label>
        <input bind:value={camera} class="p-1 rounded-lg mb-2" type="text" name="cameras" id="director">

        {#if cameras.length > 0}
            <div class="flex flex-wrap">

                {#each cameras as camera, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{camera}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeDirector(index, "c")}>X</button>
                    </div>
                {/each}

            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushDirector(camera, "c")} />

        <label class="text-cinema-text" for="genres">Genres</label>
        <input bind:value={genre} class="p-1 rounded-lg mb-2" type="text" name="genres" id="genres">

        {#if genres.length > 0}
            <div class="flex flex-wrap">

                {#each genres as genre, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{genre}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeDirector(index, "g")}>X</button>
                    </div>
                {/each}

            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushDirector(genre, "g")} />

    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Release date</legend>
        <input type="date" name="releaseDate">
    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Other</legend>
        <label class="text-cinema-text" for="country">Countries</label>
        <input bind:value={country} class="p-1 rounded-lg mb-2" type="text" name="countries" id="country">

        {#if countries.length > 0}
            <div class="flex flex-wrap">

                {#each countries as country, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{country}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeDirector(index, "country")}>X</button>
                    </div>
                {/each}

            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushDirector(country, "country")} />


        <label class="text-cinema-text" for="duration">Duration(min)</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="duration" id="duration">

        <label class="text-cinema-text" for="description">Description</label>
        <textarea class="p-1 rounded-lg mb-2" name="description" id="description" cols="10" rows="5"></textarea>

        <label class="text-cinema-text" for="actors">Actors</label>
        <input bind:value={actor} class="p-1 rounded-lg mb-2" type="text" name="actors" id="actors">

        {#if actors.length > 0}
            <div class="flex flex-wrap">

                {#each actors as actor, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{actor}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeDirector(index, "a")}>X</button>
                    </div>
                {/each}

            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushDirector(actor, "a")} />



    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Images</legend>
        <label class="text-cinema-text" for="poster">Poster</label>
        <input type="file" name="poster" id="poster">

        <label class="text-cinema-text" for="image">Images</label>

        <input type="file" on:change={pushImages} multiple name="images" id="image">
        {#if previews.length > 0}
            <div class="flex w-full h-64 whitespace-nowrap overflow-x-auto">
                {#each previews as preview, index}
                    <div class="flex shrink-0 relative">
                        <button type="button" class="text-cinema-text font-bold text-2xl absolute"
                                on:click={() => deleteImage(index)}>X</button>
                        <img class="object-cover rounded-lg w-full h-auto" src={preview} alt="Preview">
                    </div>
                {/each}
            </div>
        {/if}


    </fieldset>

    <input type="submit" value="Send" class="bg-cinema-highlight text-cinema-text">

</form>