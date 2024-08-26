<script>

    let title = "";
    let content = "";
    let preview;
    let message = "";

    async function sendArticle() {

        const formData = new FormData();

        formData.append('title', title);
        formData.append('content', content);

        console.log(preview[0])

        if (preview && preview.length > 0) {
            formData.append('preview', preview[0], preview[0].name);
        } else { console.log("No image"); }

        try {

            const response = await fetch("http://localhost:8000/add-article", {
                method: 'POST',
                body: formData,

            })

            const data = await response.json();

            if (data.success) {
                message = data.message;
            }


        } catch (e) {
            console.error(e);
            message = e;
        }


    }


</script>
<form on:submit|preventDefault={sendArticle} class="flex flex-col w-full" enctype="multipart/form-data">

    <h2 class="text-cinema-text">{message}</h2>
    <fieldset class="flex flex-col">
        <legend class="flex font-bold text-cinema-text text-xl">Title</legend>
        <label class="text-cinema-text" for="title">Enter Title</label>
        <input bind:value={title} class="p-1 rounded-lg mb-2" type="text" name="title" id="title">
    </fieldset>
    <fieldset class="flex flex-col">
        <legend class="flex font-bold text-cinema-text text-xl">Content</legend>
        <label class="text-cinema-text" for="content ">Enter Content</label>
        <textarea bind:value={content} rows="5" class="p-1 rounded-lg mb-2" name="content"
                  id="content"></textarea>
    </fieldset>
    <fieldset class="flex flex-col">
        <label class="text-cinema-text" for="preview">Image Preview</label>
        <input bind:files={preview} type="file" name="preview" id="preview">
    </fieldset>

    <input type="submit" value="Send" class="bg-cinema-highlight text-cinema-text p-2">




</form>