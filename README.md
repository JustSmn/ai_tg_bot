GOLANG AI BOT                      
Version: 1.0.0 (2025)                 
Created by: JustSMN

Welcome to my tg bot!  
This Telegram bot generates images based on your text prompts using a powerful free AI model.  
Just type your idea in English, and the bot will create a unique image for you!


Features:
- Text-to-Image Generation: Describe your idea in English, and the bot will turn it into an image. (unfortunately, sometimes AI can be slow) 
- Powered by Hugging Face: Uses the `stabilityai/stable-diffusion-2-1` model for high-quality image generation.
- Simple and Easy to Use: No complicated commands—just type and get your image!


How to Use:
1. Set the environment variables according to the bot and ai tokens:\
`PS C:\Programming\ai-tg-bot> $env:TELEGRAM_TOKEN="tgBotToken"` \
`PS C:\Programming\ai-tg-bot> $env:HUGGINGFACE_TOKEN="AiToken"` \
(For windows) 

3. Run main.go in terminal:\
               `PS C:\Programming\ai-tg-bot> go run main.go`
4. Go to telegram `@ai_golang_bot` 
5. Start the bot by typing `/start`.
6. Write your prompt in English. For example:
   - "A futuristic city at night"
   - "A cute cat wearing a hat"
   - "A magical forest with glowing mushrooms"
7. Wait for a while and bot generates your image.
8. Enjoy your unique creation! 


Technical Details:
- Model: `stabilityai/stable-diffusion-2-1` (via Hugging Face).
- Language: English only (for now).
- Version: 1.0.0 (2025 release).
- Creator: JustSMN ъ
- Contact in telegram: @Just_Semen228
- Special Thanks: I know you won't see this, but I still love you.
