<script>
    import Button from "$lib/components/Button.svelte";

    let message = "";
    let name = '';
    let year = '';
    let director = "";
    let releaseDate = '';
    let duration = '';
    let description = '';
    let writer = "";
    let producer = "";
    let editor = "";
    let camera = "";
    let genre = "";
    let country = "";
    let actor = "";
    let keyword = "";
    let budget = "";
    let language = "";
    let tmdbId = "";

    let directors = [];
    let writers = [];
    let producers = [];
    let editors = [];
    let cameras = [];
    let genres = [];
    let countries = [];
    let images = []; //files
    let actors = [];
    let previews = [];
    let keywords = [];
    let poster; // file


    function pushItem(person, type) {
        event.preventDefault();

        switch (type) {
            case "d":
                directors = [...directors, person];
                director = "";
                break;
            case "w":
                writers = [...writers, person];
                writer = "";
                break;
            case "p":
                producers = [...producers, person];
                producer = "";
                break;
            case "e":
                editors = [...editors, person];
                editor = "";
                break;
            case "c":
                cameras = [...cameras, person];
                camera = "";
                break;
            case "g":
                genres = [...genres, person];
                genre = "";
                break;
            case "a":
                actors = [...actors, person];
                actor = "";
                break;
            case "country":
                countries = [...countries, person];
                country = "";
                break;
            case "keyword":
                keywords = [...keywords, person];
                keyword = "";
                break;
            default:
                console.error("Error: Invalid type");
                break;
        }

    }

    function removeItem(index, type, event) {
        event.preventDefault();

        switch (type) {
            case "d":
                directors = directors.filter((_, i) => i !== index);
                break;
            case "w":
                writers = writers.filter((_, i) => i !== index);
                break;
            case "p":
                producers = producers.filter((_, i) => i !== index);
                break;
            case "e":
                editors = editors.filter((_, i) => i !== index);
                break;
            case "c":
                cameras = cameras.filter((_, i) => i !== index);
                break;
            case "g":
                genres = genres.filter((_, i) => i !== index);
                break;
            case "a":
                actors = actors.filter((_, i) => i !== index);
                break;
            case "country":
                countries = countries.filter((_, i) => i !== index);
                break;
            case "keyword":
                keywords = keywords.filter((_, i) => i !== index);
                break;
            default:
                console.error("Error: Invalid type");
                break;
        }
    }

    function pushImages(event) {
        const input = event.target;
        if (input && input.files.length > 0) {
            const newImages = Array.from(input.files).filter(file => !images.includes(file.name));
            const newPreviews = newImages.map(file => URL.createObjectURL(file));
            images = [...images, ...newImages];
            previews = [...previews, ...newPreviews];
            console.log(images);
        }
    }

    function deleteImage(index) {
        images = images.filter((_, i) => i !== index);
        previews = previews.filter((_, i) => i !== index);
    }


   async function newData() {

        const formData = new FormData();
        formData.append('name', name);
        formData.append('year', year);
        formData.append('budget', budget);
        formData.append('language', language);
        directors.forEach(director => {formData.append('directors', director);})
        writers.forEach(    writer  => {formData.append('writers', writer)});
        producers.forEach(  producer => {formData.append('producers', producer)});
        editors.forEach(    editor  => {formData.append('editors', editor)});
        cameras.forEach(    camera  => {formData.append('cameras', camera)});
        genres.forEach(     genre   => {formData.append('genres', genre)});
        actors.forEach(     actor   => {formData.append('actors', actor)});
        countries.forEach(  country => {formData.append('countries', country)});
        keywords.forEach(   keyword => {formData.append('keywords', keyword)});

        formData.append('releaseDate', releaseDate);
        formData.append('duration', duration);
        formData.append('description', description);
        formData.append('tmdbId', tmdbId);


       if (poster && poster.length > 0) {
           formData.append('poster', poster[0], poster[0].name);
       }

       if (images) {
           for (let i = 0; i < images.length; i++) {
               console.log("Processing", images[i]);
               formData.append('images', images[i]);
           }
       }

       try {
           const response = await fetch("http://localhost:8000/add-movie", {
               method: 'POST',
               body: formData,
           })

           const data = await response.json();

           if (data.success) {

               message = data.message;

           }

       } catch (e) {
           message = e;
       }



        console.log(formData)
    }
</script>

<h2 class="text-2xl text-white">{#if message}{message}{/if}</h2>
<form on:submit|preventDefault={newData} class="flex flex-col bg-cinema-secondary w-full" enctype="multipart/form-data">
    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Basic Info</legend>
        <label class="text-cinema-text" for="movie-name">Movie Name</label>
        <input bind:value={name} class="p-1 rounded-lg mb-2" type="text" name="movie-name" id="movie-name">
        <label class="text-cinema-text" for="year">Year</label>
        <input bind:value={year} class="p-1 rounded-lg mb-2" type="text" name="year" id="year">
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
                                on:click={() => removeItem(index, "d")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(director, "d")} />

        <label class="text-cinema-text" for="writers">Writers</label>
        <input bind:value={writer} class="p-1 rounded-lg mb-2" type="text" name="writers" id="writers">

        {#if writers.length > 0}
            <div class="flex flex-wrap">
                {#each writers as writer, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{writer}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeItem(index, "w")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(writer, "w")} />

        <label class="text-cinema-text" for="producers">Producers</label>
        <input bind:value={producer} class="p-1 rounded-lg mb-2" type="text" name="producers" id="producers">

        {#if producers.length > 0}
            <div class="flex flex-wrap">
                {#each producers as producer, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{producer}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeItem(index, "p")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(producer, "p")} />

        <label class="text-cinema-text" for="editors">Editors</label>
        <input bind:value={editor} class="p-1 rounded-lg mb-2" type="text" name="editors" id="editors">

        {#if editors.length > 0}
            <div class="flex flex-wrap">
                {#each editors as editor, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{editor}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeItem(index, "e")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(editor, "e")} />

        <label class="text-cinema-text" for="cameras">Cameras</label>
        <input bind:value={camera} class="p-1 rounded-lg mb-2" type="text" name="cameras" id="cameras">

        {#if cameras.length > 0}
            <div class="flex flex-wrap">
                {#each cameras as camera, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{camera}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeItem(index, "c")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(camera, "c")} />

        <label class="text-cinema-text" for="genres">Genres</label>
        <input bind:value={genre} class="p-1 rounded-lg mb-2" type="text" name="genres" id="genres">

        {#if genres.length > 0}
            <div class="flex flex-wrap">
                {#each genres as genre, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{genre}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeItem(index, "g")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(genre, "g")} />

    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Release date</legend>
        <input bind:value={releaseDate} type="date" name="releaseDate">
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
                                on:click={() => removeItem(index, "country")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(country, "country")} />

        <label class="text-cinema-text" for="duration">Duration(min)</label>
        <input bind:value={duration} class="p-1 rounded-lg mb-2" type="text" name="duration" id="duration">

        <label class="text-cinema-text" for="budget">Budget</label>
        <input bind:value={budget} class="p-1 rounded-lg mb-2" type="text" name="budget" id="budget">

        <label class="text-cinema-text" for="language">Language</label>
        <input bind:value={language} class="p-1 rounded-lg mb-2" type="text" name="language" id="language">

        <label class="text-cinema-text" for="tmdbId">TMDB ID</label>
        <input bind:value={tmdbId} class="p-1 rounded-lg mb-2" type="text" name="tmdbId" id="tmdbId">

        <label class="text-cinema-text" for="description">Description</label>
        <textarea bind:value={description} class="p-1 rounded-lg mb-2" name="description" id="description"
                  cols="10" rows="5">
        </textarea>

        <label class="text-cinema-text" for="actors">Actors</label>
        <input bind:value={actor} class="p-1 rounded-lg mb-2" type="text" name="actors" id="actors">

        {#if actors.length > 0}
            <div class="flex flex-wrap">
                {#each actors as actor, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{actor}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeItem(index, "a")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(actor, "a")} />

        <label class="text-cinema-text" for="actors">Keywords</label>
        <input bind:value={keyword} class="p-1 rounded-lg mb-2" type="text" name="actors" id="actors">

        {#if keywords.length > 0}
            <div class="flex flex-wrap">
                {#each keywords as keyword, index}
                    <div class="flex justify-between items-center pl-2">
                        <span class="text-cinema-text">{keyword}</span>
                        <button type="button" class="text-cinema-text font-bold text-xl ml-2"
                                on:click={() => removeItem(index, "keyword")}>X</button>
                    </div>
                {/each}
            </div>
        {/if}
        <Button padding="p-2" name="Push" on:click={() => pushItem(keyword, "keyword")} />


    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Images</legend>
        <label class="text-cinema-text" for="poster">Poster</label>
        <input bind:files={poster} type="file" name="poster" id="poster">

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
