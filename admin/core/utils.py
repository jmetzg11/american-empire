import os
import time
from django.conf import settings
from supabase import create_client, Client

def save_uploaded_photo(uploaded_file, event_id):
    filename = f"{int(time.time())}_{uploaded_file.name}"
    
    if os.environ.get('ENVIRONMENT') == 'production':
        path = f"{event_id}/{filename}"
        
        supabase_url = os.environ.get('SUPABASE_URL')
        supabase_key = os.environ.get('SUPABASE_SERVICE_ROLE_KEY')
        
        if not supabase_url or not supabase_key:
            raise ValueError("Supabase credentials not configured")
            
        supabase: Client = create_client(supabase_url, supabase_key)
        
        try:
            uploaded_file.seek(0)  # Reset file pointer
            result = supabase.storage.from_("photos").upload(path, uploaded_file.read())
            
            if result.get('error'):
                raise Exception(f"Supabase upload error: {result['error']}")
                
            return path
        except Exception as e:
            raise Exception(f"Failed to upload to Supabase: {str(e)}")
    else:
        dir_path = f"data/photos/{event_id}"
        full_dir_path = os.path.join(settings.BASE_DIR.parent, dir_path)
        
        os.makedirs(full_dir_path, exist_ok=True)
        
        file_path = os.path.join(full_dir_path, filename)
        
        with open(file_path, 'wb+') as destination:
            for chunk in uploaded_file.chunks():
                destination.write(chunk)
        
        return f"{event_id}/{filename}"
