<script>
    export let title = "Article Name";
    export let poster = "";
    export let content = "This is the article's content";
</script>



<div class="flex flex-col shrink-0 w-1/2 rounded-lg p-2">
    <a href="#" class="flex">
        <img class="object-cover rounded-lg w-full h-auto" src={poster} alt={title} />
    </a>
    <a href="#" class="text-cinema-text line-clamp-2 py-4 ">{title}</a>
    <a href="#" class="flex text-cinema-text py-4 ">
        <p class="line-clamp-1">{content}</p>
    </a>
</div>