from django.db import models

class Event(models.Model):
    id = models.AutoField(primary_key=True)
    title = models.CharField(max_length=255)
    date = models.DateTimeField(db_index=True)
    country = models.CharField(max_length=255, db_index=True)
    description = models.TextField()
    created_at = models.DateTimeField(auto_now_add=True)
    updated_at = models.DateTimeField(auto_now=True)
    active = models.DateTimeField(null=True, blank=True)
    flagged = models.BooleanField(default=False)

    class Meta:
        db_table = 'events'

    def __str__(self):
        return self.title

class Source(models.Model):
    id = models.AutoField(primary_key=True)
    event = models.ForeignKey(Event, on_delete=models.CASCADE, related_name='sources')
    name = models.CharField(max_length=255)
    url = models.URLField()

    class Meta:
        db_table = 'sources'

    def __str__(self):
        return self.name

class Media(models.Model):
    id = models.AutoField(primary_key=True)
    event = models.ForeignKey(Event, on_delete=models.CASCADE, related_name='medias')
    type = models.CharField(max_length=100, blank=True)
    url = models.URLField(blank=True)
    path = models.CharField(max_length=500, blank=True)
    caption = models.CharField(max_length=500, blank=True)

    class Meta:
        db_table = 'media'

    def __str__(self):
        return f"{self.type} for {self.event.title}"

class Tag(models.Model):
    id = models.AutoField(primary_key=True)
    name = models.CharField(max_length=255, unique=True)
    events = models.ManyToManyField(Event, through='EventTag', related_name='tags')

    class Meta:
        db_table = 'tags'

    def __str__(self):
        return self.name

class Book(models.Model):
    id = models.AutoField(primary_key=True)
    title = models.CharField(max_length=255)
    author = models.CharField(max_length=255)
    link = models.URLField()
    events = models.ManyToManyField(Event, through='BookEvent', related_name='books')
    tags = models.ManyToManyField(Tag, through='BookTag', related_name='books')

    class Meta:
        db_table = 'books'

    def __str__(self):
        return self.title

class EventTag(models.Model):
    event = models.ForeignKey(Event, on_delete=models.CASCADE)
    tag = models.ForeignKey(Tag, on_delete=models.CASCADE)

    class Meta:
        db_table = 'event_tags'

class BookEvent(models.Model):
    book = models.ForeignKey(Book, on_delete=models.CASCADE)
    event = models.ForeignKey(Event, on_delete=models.CASCADE)

    class Meta:
        db_table = 'book_events'

class BookTag(models.Model):
    book = models.ForeignKey(Book, on_delete=models.CASCADE)
    tag = models.ForeignKey(Tag, on_delete=models.CASCADE)

    class Meta:
        db_table = 'book_tags'