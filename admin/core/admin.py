from django.contrib import admin
from .models import Event, Source, Media, Tag, Book

class SourceInline(admin.TabularInline):
    model = Source
    extra = 1

class MediaInline(admin.TabularInline):
    model = Media
    extra = 1

@admin.register(Event)
class EventAdmin(admin.ModelAdmin):
    list_display = ['title', 'date', 'country', 'active', 'flagged', 'created_at']
    list_filter = ['country', 'flagged', 'date']
    search_fields = ['title', 'description', 'country']
    inlines = [SourceInline, MediaInline]

@admin.register(Tag)
class TagAdmin(admin.ModelAdmin):
    list_display = ['name']
    search_fields = ['name']

@admin.register(Book)
class BookAdmin(admin.ModelAdmin):
    list_display = ['title', 'author']
    search_fields = ['title', 'author']

@admin.register(Source)
class SourceAdmin(admin.ModelAdmin):
    list_display = ['name', 'event', 'url']
    list_filter = ['event']

@admin.register(Media)
class MediaAdmin(admin.ModelAdmin):
    list_display = ['type', 'event', 'caption', 'url', 'path']
    list_filter = ['type', 'event']
