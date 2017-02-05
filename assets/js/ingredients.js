function addIngredient(name, id) {
    table = $('#ingredients_table');
    console.log(id);
    close = $('<span>', { class: 'remove' }).html('[✗]');
    $('<li>', { class: 'ingredient-elem', 'data-id': id})
        .append($('<span>', { class: 'name' }).html(name.trim()))
        .append(close)
        .appendTo(table);
    close.click(deleteCallback);
}

// Initializes form to submit new ingredient
function initForm() {
    $('#new_ingredient_form').submit((event) => {
        event.preventDefault();
        req_url = '/ingredient/new?' + $('#new_ingredient_form').serialize();
        post = $.post(req_url);

        post.fail(formRejected);
        post.done(formAccepted);
    });
};

// Action for failed post request
function formRejected(data) {
    console.log(data);
}

// Action for accepted post request
function formAccepted(data) {
    data = JSON.parse(data);
    nameInput = $('#new_ingredient_form input[name=name]');
    addIngredient(nameInput.val(), data.id);
    nameInput.val('');
    $('html').scrollTop($(document).height());
}

// Callbacks to delete ingredient
function initDeleteActions() {
    $('span.remove').each(function(index) {
        $(this).click(deleteCallback);
    });
}

function deleteCallback() {
    id = $(this).parent().data('id');
    request = $.ajax({
        url: '/ingredient/' + id,
        type: 'DELETE',
        success: deleteAccepted(id),
        error: deleteRejected
    });
    request.fail(deleteRejected);
}

function deleteRejected(data) {
    console.log(data);
}

function deleteAccepted(id) {
    return function(data) {
        $('li.ingredient-elem[data-id="'+id+'"]').remove();
    };
}

// Entry point on page load
$( () => {
    initForm();
    initDeleteActions();
});
