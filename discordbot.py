import discord, subprocess

token = "<discord-token>"
bot = discord.Client()

@bot.event
async def on_ready():
        print(f"bot ID : {bot.user.id}")
        print("SampleDiscordBot is in " + str(guild_count) + " guilds.")

@bot.event
async def on_message(message):
        if str(bot.user.id) in message.content:
                print('bot is mentioned')
                cmd = message.content.split()
                print(cmd[1])
                await message.channel.send("thanks for the mention bro")
                if cmd[1] == "hello":
                        await message.channel.send("hey we will do good work together")
                if cmd[1] == "uname":
                        cmdcall = subprocess.run(["uname", "-a"], stdout=subprocess.PIPE, text=True)
                        await message.channel.send(cmdcall.stdout)
                if cmd[1] == "ifconfig":
                        cmdcall = subprocess.run(["ifconfig"], stdout=subprocess.PIPE, text=True)
                        await message.channel.send(cmdcall.stdout)
                if cmd[1] == "whoami":
                        cmdcall = subprocess.run(["whoami"], stdout=subprocess.PIPE, text=True)
                        await message.channel.send(cmdcall.stdout)

bot.run(token)
