import os
from django.db import models
from django.utils.html import format_html
from django.utils.safestring import mark_safe

class Event(models.Model):
    id = models.AutoField(primary_key=True)
    title = models.CharField(max_length=255)
    date = models.DateField(db_index=True)
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
    MEDIA_TYPE_CHOICES = [
        ('photo', 'Photo'),
        ('youtube', 'YouTube'),
    ]
    
    id = models.AutoField(primary_key=True)
    event = models.ForeignKey(Event, on_delete=models.CASCADE, related_name='medias')
    type = models.CharField(max_length=100, choices=MEDIA_TYPE_CHOICES, blank=True)
    url = models.CharField(max_length=500, blank=True, help_text="For YouTube: enter video ID (e.g., LDYwIOWbeRU) or full URL")
    path = models.CharField(max_length=500, blank=True)
    caption = models.CharField(max_length=500, blank=True)

    class Meta:
        db_table = 'media'

    def __str__(self):
        return f"{self.type} for {self.event.title}"
    
    def image_preview(self):
        if self.path and self.type and self.type.lower() == 'photo':
            if os.environ.get('ENVIRONMENT') == 'production':
                # Production: Use Supabase URL
                supabase_url = os.environ.get('SUPABASE_URL', '')
                if supabase_url:
                    image_url = f"{supabase_url}/storage/v1/object/public/photos/{self.path}"
                else:
                    return "Supabase not configured"
            else:
                # Development: Use Django media serving
                image_url = f"/media/photos/{self.path}"
            return format_html('<img src="{}" style="max-width: 200px; max-height: 200px;" />', image_url)
        return "No image"
    
    image_preview.short_description = "Preview"
    
    def delete(self, *args, **kwargs):
        if self.path and self.type and self.type.lower() == 'photo':
            if os.environ.get('ENVIRONMENT') == 'production':
                try:
                    from supabase import create_client
                    supabase_url = os.environ.get('SUPABASE_URL')
                    supabase_key = os.environ.get('SUPABASE_SERVICE_ROLE_KEY')
                    
                    if supabase_url and supabase_key:
                        supabase = create_client(supabase_url, supabase_key)
                        supabase.storage.from_("photos").remove([self.path])
                except Exception as e:
                    print(f"Warning: Failed to delete photo from Supabase: {e}")
            else:
                # Development: Delete local file
                try:
                    from django.conf import settings
                    file_path = os.path.join(settings.BASE_DIR.parent, 'data', 'photos', self.path)
                    if os.path.exists(file_path):
                        os.remove(file_path)
                except Exception as e:
                    # Log the error but don't prevent deletion of the database record
                    print(f"Warning: Failed to delete local photo file: {e}")
        
        super().delete(*args, **kwargs)

class Tag(models.Model):
    id = models.AutoField(primary_key=True)
    name = models.CharField(max_length=255, unique=True)
    events = models.ManyToManyField(Event, db_table='event_tags', related_name='tags')

    class Meta:
        db_table = 'tags'

    def __str__(self):
        return self.name

class Book(models.Model):
    id = models.AutoField(primary_key=True)
    title = models.CharField(max_length=255)
    author = models.CharField(max_length=255)
    link = models.URLField()
    events = models.ManyToManyField(Event, db_table='book_events', related_name='books')
    tags = models.ManyToManyField(Tag, db_table='book_tags', related_name='books')

    class Meta:
        db_table = 'books'

    def __str__(self):
        return self.title

