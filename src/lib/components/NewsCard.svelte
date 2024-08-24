<script>
    export let title = "Article Name";
    export let poster = "";
    export let content = "This is the article's content";
</script>



<div class="flex flex-col shrink-0 w-1/2 rounded-lg p-2">
    <a href="image" class="flex">
        <img class="object-cover rounded-lg w-full h-auto" src={poster} alt={title} />
    </a>
    <a href="image" class="text-cinema-text line-clamp-2 py-4 ">{title}</a>
    <a href="image" class="flex text-cinema-text py-4 ">
        <p class="line-clamp-1">{content}</p>
    </a>
</div>