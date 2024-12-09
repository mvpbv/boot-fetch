import numpy as np
import matplotlib.pyplot as plt
from matplotlib.dates import DateFormatter
from datetime import timedelta
import pandas as pd
import os
import glob

dir_path = 'websrc'
all_files = glob.glob(os.path.join(dir_path, "*.csv"))
li = []

for filename in all_files:
    df = pd.read_csv(filename, index_col=None, header=0)
    df['user'] = filename.split('/')[-1].split('.')[0]
    li.append(df)
users = pd.concat(li, axis=0, ignore_index=True)
users['Time'] = pd.to_datetime(users['Time'])
users.value_counts('user')


def get_total_xp(level):
    return (level * (level-1) * 40) + (400 * level)
for user in users['user'].unique():
    ti_min = users.where(users['user'] == user)['Time'].min()
    ti_max = users.where(users['user'] == user)['Time'].max()
    xp_max = users.where(users['user'] == user)['Total_Xp'].max()
    xp_min = users.where(users['user'] == user)['Total_Xp'].min()
    t_delta = ti_max - ti_min
    
    xp_delta = xp_max - xp_min
    seconds = t_delta.total_seconds()

    if seconds == 0:
        users = users.loc[users['user'] != user]
        continue
    xp_per_day = xp_delta / (seconds / (3600 * 24))
    
users['Time'] = pd.to_datetime(users['Time'])
users['Total_Xp'] = users['Level'].apply(get_total_xp)
users['Total_Xp'] = users['Xp'] + users['Total_Xp']
first_users = {}
for user in users['user'].unique():
    first = users.where(users['user'] == user).dropna().sort_values('Time').iloc[0]
    first_users[user] = (first['Total_Xp'], first['Time'])
fig, ax = plt.subplots(figsize=(18, 18))
users.groupby('user')[['Total_Xp', 'Time']].plot(kind='line',x='Total_Xp', y='Time', legend=False, ax=ax, alpha=0.8)

plt.xlim(get_total_xp(10), get_total_xp(100))
x_ticks = [get_total_xp(i) for i in range(0, 101, 10)]
role_names = ['newb', 'pupil', 'acolyte', 'disciple', 'scholar', 'sorceror', 'sage', 'archsage', 'necromancer', 'archmage']
date_form = DateFormatter("%m-%d")
x_ticks.remove(7600)

ax.yaxis.set_major_formatter(date_form)
ax.set_facecolor('#1f2335')
fig.set_facecolor('#545c7e')
names = sorted(users['user'].unique())
line_colors = ['#ff007c', '#ff6731', '#69ff46', '#39ff14', '#0062FF','#b4f9f8', '#d875ff','#ff9e64', '#7dcfff', '#0ff0fc', '#4fd6be',
               '#8d4be0', '#ccff00','#567cff','#d875ff', '#7aa2f7','#000000', '#ff757f', '#567cff','#c3e88d','#ccff00','#c3e88d',
               '#9d7cd8','#ff6731', '#3d69a1','#c53b53' , '#41a6b5', '#3a90e5', '#394b80', '#ffc777', '#1562c1', '#ffffff']
user_attr = {}


for first_user in first_users:
    """
    Create offsets and specify the color of the text manually if your graph gets smushed with additional users or time.
    """
    ax.annotate(xy=(first_users[first_user][0] - 5000, first_users[first_user][1] - timedelta(hours=8)), color=user_attr[first_user], text=first_user,)
plt.rcParams.update({'font.size': 11})
plt.xlabel('Total Xp')
plt.title('The Road to Archmage')
plt.xticks(x_ticks, role_names)
v_lines_five = [(get_total_xp(i) + get_total_xp(i + 10))/ 2  for i in range(50, 90, 10)]
v_lines_five = list(filter(lambda x: x not in x_ticks, v_lines_five))
v_lines_three = [get_total_xp(i) for i in range(93, 100, 4)]
y_lim_low, y_lim_high = ax.get_ylim()[0], ax.get_ylim()[1]
plt.vlines(x_ticks, y_lim_low, y_lim_high, color='#545c7e', linestyles='dashed', alpha=0.5)
plt.vlines(v_lines_five, y_lim_low, y_lim_high, color='#545c7e', linestyles='dashed', alpha=0.3)
plt.vlines(v_lines_three, y_lim_low, y_lim_high, color='#545c7e', linestyles='dashed', alpha=0.3)
plt.xlim(get_total_xp(0), get_total_xp(101))


fig = plt.gcf()
plt.show()


fig.savefig('user_plot.svg', format='svg')
fig.savefig('boot-fetch.png', format='png')
