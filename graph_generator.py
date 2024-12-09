#import numpy as np
import pandas as pd 
import matplotlib.pyplot as plt
from matplotlib.dates import DateFormatter
from datetime import timedelta
import sqlalchemy as db

def lvl_to_xp(lvl):
    return (lvl * (lvl - 1) * 40) + (lvl * 400)

def connect_Load():
    engine = db.create_engine('postgresql://braedenvallejos@localhost:5432/vinny_fetch')
    conn = engine.connect()
    metadata = db.MetaData()
    df = pd.read_sql('SELECT progress.total_xp, users.discord_name, progress.time FROM progress JOIN users ON users.id = progress.user_id', conn)
    df.sort_values(by='time', inplace=True)
    return df
def inital_time(df):
    init_time = {}
    for user in df['discord_name'].unique():
        first = df[df['discord_name'] == user].iloc[0]
        init_time[user] = (first['total_xp'], first['time'])
    return init_time
def set_colors(ax, fig):
    ax.set_facecolor('1f2335')
    fig.set_facecolor('545c7e')
def set_x_ticks(ax):
    x_ticks = [lvl_to_xp(i) for i in range(10, 101, 10)]
    x_ticks.remove(7600)
    role_names = ['newb', 'pupil', 'acolyte', 'disciple', 'scholar', 'sorceror', 'sage', 'archsage', 'necromancer', 'archwizzer']
    plt.xticks(x_ticks, role_names)
    plot_vlines(x_ticks, ax)

def plot_vlines(x_ticks, ax):
    y_lim_low, y_lim_high = ax.get_ylim()[0], ax.get_ylim()[1]
    v_lines_five = [(lvl_to_xp(i) + lvl_to_xp(i + 10))/ 2 for i in range(50, 101, 10)]
    plt.vlines(x_ticks, y_lim_low, y_lim_high, color='#545c7e', linestyles='dotted', alpha=0.5)
    plt.vlines(v_lines_five, y_lim_low, y_lim_high, color='545c73', linestyles='dotted', alpha=0.5)

def create_plot(x, y, df):
    fig, ax = plt.subplots(figsize=(x, y))
    df.groupby('discord_name').plot(x='time', y='total_xp', ax=ax)
    plt.xlim(lvl_to_xp(10), lvl_to_xp(100))
    
    date_fmt = DateFormatter('%m-%d')
 
    ax.yaxis.set_major_formatter(date_fmt)
   




