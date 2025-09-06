from django.contrib import admin
from django import forms
from .models import Event, Source, Media, Tag, Book
from .utils import save_uploaded_photo

class SourceInline(admin.TabularInline):
    model = Source
    extra = 1

class MediaAdminForm(forms.ModelForm):
    upload_file = forms.FileField(required=False, help_text="Upload a photo file")
    
    class Meta:
        model = Media
        fields = '__all__'
    
    def save(self, commit=True):
        instance = super().save(commit=False)
        uploaded_file = self.cleaned_data.get('upload_file')
        
        if uploaded_file and instance.event_id:
            # Save the uploaded file and set the path
            try:
                path = save_uploaded_photo(uploaded_file, instance.event_id)
                instance.path = path
                instance.type = 'photo'
            except Exception as e:
                raise forms.ValidationError(f"Failed to upload file: {str(e)}")
        
        if commit:
            instance.save()
        return instance

class MediaInline(admin.TabularInline):
    model = Media
    form = MediaAdminForm
    extra = 1
    readonly_fields = ['image_preview']
    can_delete = True
    show_change_link = True
    fields = ['type', 'url', 'path', 'caption', 'upload_file', 'image_preview']

@admin.register(Event)
class EventAdmin(admin.ModelAdmin):
    list_display = ['title', 'date', 'country', 'active', 'flagged', 'created_at']
    list_filter = ['country', 'flagged', 'date', 'active']
    search_fields = ['title', 'description', 'country']
    inlines = [SourceInline, MediaInline]
    readonly_fields = ['get_tags']
    
    def get_tags(self, obj):
        return ', '.join([tag.name for tag in obj.tags.all()])
    get_tags.short_description = 'Tags'

@admin.register(Tag)
class TagAdmin(admin.ModelAdmin):
    list_display = ['name']
    search_fields = ['name']
    filter_horizontal = ['events']

@admin.register(Book)
class BookAdmin(admin.ModelAdmin):
    list_display = ['title', 'author']
    search_fields = ['title', 'author']
    filter_horizontal = ['events']

@admin.register(Source)
class SourceAdmin(admin.ModelAdmin):
    list_display = ['name', 'event', 'url']
    list_filter = ['event']

@admin.register(Media)
class MediaAdmin(admin.ModelAdmin):
    form = MediaAdminForm
    list_display = ['type', 'event', 'caption', 'url', 'path', 'image_preview']
    list_filter = ['type', 'event']
    readonly_fields = ['image_preview']
    fields = ['event', 'type', 'url', 'path', 'caption', 'upload_file', 'image_preview']
