import discord, subprocess

token = "<discord-token>"
bot = discord.Client()

@bot.event
async def on_ready():
        print(f"bot ID : {bot.user.id}")
        print("SampleDiscordBot is in " + str(guild_count) + " guilds.")

@bot.event
async def on_message(message):
        print(message.content)
        if str(bot.user.id) in message.content:
                print('bot is mentioned')
                cmd = message.content.split()
                await message.channel.send("thanks for the mention bro")
                if cmd[1] == "cmd":
                        command = []
                        for i in range(len(cmd)):
                                if i > 1:
                                        command.append(cmd[i])
                        print(command)
                        cmdcall = subprocess.run(command, stdout=subprocess.PIPE, text=True)
                        await message.channel.send(cmdcall.stdout)

bot.run(token)
