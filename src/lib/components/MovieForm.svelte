<script>
    import Button from "$lib/components/Button.svelte";
    import {json} from "@sveltejs/kit";

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
    let poster;

    function pushItem(person, type) {
        event.preventDefault();

        switch (type) {
            case "d":
                directors.push(person)
                directors = directors;
                break;
            case "w":
                writers = [...writers, person];
                break;
            case "p":
                producers = [...producers, person];
                break;
            case "e":
                editors = [...editors, person];
                break;
            case "c":
                cameras = [...cameras, person];
                break;
            case "g":
                genres = [...genres, person];
                break;
            case "a":
                actors = [...actors, person];
                break;
            case "country":
                countries = [...countries, person];
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
            default:
                console.error("Error: Invalid type");
                break;
        }
    }

    function pushImages(event) {
        const input = event.target;
        if (input && input.files.length > 0) {
            const newImages = Array.from(input.files);
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

    async function sendData(event) {
        event.preventDefault();

        const formData = new FormData();

        if (images) {
            for (let i = 0; i < images.length; i++) {
                formData.append('images', images[i]);
            }
        }

        formData.append("directors", JSON.stringify(directors));


        // Debug logs to check the form data before sending it
        console.log("Directors:", directors);


        try {
            const response = await fetch("http://localhost:8000/add-movie", {
                method: "POST",
                body: formData,
            });

            const data = await response.json();

            if (data.success) {
                console.log("Movie added successfully:", (message = data.message));
            } else {
                console.error("Error adding movie:", (message = data.message));
            }
        } catch (e) {
            console.error(e);
        }
    }


   async function newData() {

        const formData = new FormData();
        formData.append('name', name);
        formData.append('year', year);
        directors.forEach(director => {

            formData.append('directors', director);
        })
        formData.append('releaseDate', releaseDate);
        formData.append('duration', duration);
        formData.append('description', description);
        formData.append('poster', poster);

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

    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Release date</legend>
        <input bind:value={releaseDate} type="date" name="releaseDate">
    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Other</legend>

        <label class="text-cinema-text" for="duration">Duration(min)</label>
        <input bind:value={duration} class="p-1 rounded-lg mb-2" type="text" name="duration" id="duration">

        <label class="text-cinema-text" for="description">Description</label>
        <textarea bind:value={description} class="p-1 rounded-lg mb-2" name="description" id="description"
                  cols="10" rows="5"></textarea>

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
