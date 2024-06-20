<script>




let images = [];
let previews = [];

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
    previews = previews.filter((_,i) => i !== index);
}
</script>



<form class="flex flex-col bg-cinema-secondary w-full" method="POST">
    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Initials</legend>
        <label class="text-cinema-text " for="name">First Name</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="name" id="name">
        <label class="text-cinema-text" for="lastName">Last Name</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="lastName" id="lastName">
    </fieldset>
    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">Biography</legend>
        <label class="text-cinema-text" for="birthday">Day of Birth</label>
        <input class="p-1 rounded-lg mb-2" type="date" name="birthday" id="birthday">
        <label class="text-cinema-text" for="gender">Gender</label>
        <select class="p-1 rounded-lg w-1/2 mb-2" id="gender" name="gender">
            <option value="male">Male</option>
            <option value="female">Female</option>
        </select>
        <label class="text-cinema-text" for="pob">Place of Birth</label>
        <input class="p-1 rounded-lg mb-2" type="text" name="pob" id="pob">

        <label class="text-cinema-text" for="biog">Biography</label>
        <textarea class="p-1 rounded-lg mb-2" name="biog" id="biog" cols="10" rows="5"></textarea>
    </fieldset>

    <fieldset class="flex flex-col p-2">
        <legend class="flex font-bold text-cinema-text text-xl">
            Media
        </legend>
        <label class="text-cinema-text " for="image">Add Images</label>
        <input type="file" on:change={pushImages} multiple name="images" id="image">

        {#if previews.length > 0}
            <div class="flex w-full h-64 whitespace-nowrap overflow-x-auto">
            {#each previews as preview, index}
                <div class="flex shrink-0">
                    <button type="button" class="text-cinema-text font-bold text-2xl absolute"
                            on:click={() => deleteImage(index)}>X</button>
                    <img class="object-cover rounded-lg w-full h-auto" src={preview} alt="Preview">
                </div>
                {/each}
            </div>
            {/if}


    </fieldset>

    <input type="submit" value="Submit">
</form>