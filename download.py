#-*- coding:utf-8 -*-
import requests
import sys
reload(sys)
sys.setdefaultencoding( "utf-8" )
from lxml import etree
import pymysql
import copy

def download_noteData():
    print "========start"
    #for i in range(1,2448):
    for i in range(2459,2482):
        url = "http://timetag.main.jp/nicoflick/nicoflick.php?req=timetag&id="+str(i)+"&no-json=1&form=1"
        r = requests.get(url)
        #print r.text
        with open("/Users/tiger/Project/noteData/noteData_"+str(i)+".txt", mode='w') as f:
            f.write(r.text)
        f.close()
        
    print "========finish"

def download_noteDataV2(i,url):
    print "========start"
    try:
        r = requests.get(url)
        #不应该是写文件，而是直接返回，并且需要时utf-8
        with open("/Users/tiger/Project/noteData/noteData_"+str(i)+".txt", mode='w') as f:
            f.write(r.text)
        f.close()
    except Exception,e:
        print e
        
        
    print "========finish"


def get_max():
    print "=========start"
    url = "http://timetag.main.jp/nicoflick/index.php"
    r = requests.get(url)
    text = r.content.decode("utf-8")
    html = etree.HTML(text)
    music_max = html.xpath("/html/body/div[3]/div/table/tbody/tr[1]/td[1]/text()")
    level_max = html.xpath("/html/body/table/tbody/tr[1]/td[1]/text()")
    return music_max[0],level_max[0]


def postSql_getData():
    print "=========start"
    url = "http://timetag.main.jp/nicoflick/index.php"
    m,l = get_max()
    #body = {
    #    'music-order': "ORDER BY id DESC LIMIT 0,"+m,
    #    'level-order': "ORDER BY id DESC LIMIT 0,"+l
    #}  
    body = {
        'music-order': "ORDER BY id DESC LIMIT 0,1",
        'level-order': "ORDER BY id DESC LIMIT 0,1"
    }    
    r = requests.post(url,data=body)
    text = r.content.decode("utf-8")
    html = etree.HTML(text)
    music_data_ = []
    level_data_ = []
    level_data = {}
    music_data = {}
    
    #/html/body/div[3]/div/table/tbody/tr
    music_lis = html.xpath("//div[3]/div/table/tbody/tr")
    #print len(music_lis)
    #//div[3]/div/table/tbody/tr[1]/td[1]
    for li in range(0,len(music_lis)):
        music_param_lis = html.xpath("//div[3]/div/table/tbody/tr["+str(li+1)+"]/td/text()")
        #print len(music_param_lis)
        music_id = music_param_lis[0]
        music_movie_url = music_param_lis[1]
        music_thumbnail_url = music_param_lis[2]
        music_title = music_param_lis[3]
        music_artist = music_param_lis[4]
        music_length = music_param_lis[5]
        music_tags = music_param_lis[6]
        music_updatetime = music_param_lis[7]
        music_createtime = music_param_lis[8]
        #etree.tostring(music_id,pretty_print=True,encoding='utf-8').decode('utf-8')
        #print etree.tostring(music_id,encoding='utf-8').decode('utf-8')
        
                
        music_data["id"] = music_id
        music_data["movie_url"] = music_movie_url
        music_data["thumbnail_url"] = music_thumbnail_url
        music_data["title"] = music_title
        music_data["artist"] = music_artist
        music_data["movie_length"] = music_length
        music_data["tags"] = music_tags
        music_data["update_time"] = music_updatetime
        music_data["create_time"] = music_createtime
        
        #deepcopy reuse map        
        music_data_.append(copy.deepcopy(music_data))
        music_data.clear()
    
    #//table/tbody/tr[1]    /html/body/table/tbody/tr[1]
    level_lis = html.xpath("/html/body/table/tbody/tr")
    for li in range(0,len(level_lis)):
        #/html/body/table/tbody/tr[1]/td[1]
        level_param_lis = html.xpath("/html/body/table/tbody/tr["+str(li+1)+"]/td/text()")
        #print len(level_param_lis)
        level_id = level_param_lis[0]
        level_movie_url = level_param_lis[1]
        level_l = level_param_lis[2]
        level_creator = level_param_lis[3]
        level_des = level_param_lis[4]
        level_speed = level_param_lis[5]
        level_note_url = "http://timetag.main.jp/nicoflick/nicoflick.php?req=timetag&id="+level_id+"&no-json=1&form=1"
        level_note = "Link"
        #download_noteDataV2(level_id,level_note_url)
        level_updatetime = level_param_lis[6]
        level_createtime = level_param_lis[7]        
        
        level_data["id"] = level_id
        level_data["movie_url"] = level_movie_url
        level_data["level"] = level_l
        level_data["creator"] = level_creator
        level_data["description"] = level_des
        level_data["speed"] = level_speed
        level_data["notes"] = level_note
        level_data["update_time"] = level_updatetime
        level_data["create_time"] = level_createtime
        
        #deepcopy reuse map
        level_data_.append(copy.deepcopy(level_data))
        level_data.clear()
    

    uploadToSql(music_data_,level_data_)

def uploadToSql(m,l):
    print "===========start"
    db = pymysql.connect("127.0.0.1","root","","flickdb")
    cursor = db.cursor()
       
    for i in m:
        try:
            #insert
            sql1 = "UPDATE music_data SET id = %d ,movie_url = '%s' ,thumbnail_url = '%s' ,title = '%s' ,artist ='%s', movie_length = '%s', tags = '%s' ,update_time = %d ,create_time = %d" % (int(i.get("id")),
                                                                                                                                                                                                i.get("movie_url"),
                                                                                                                                                                                                i.get("thumbnail_url"),
                                                                                                                                                                                                i.get("title"),
                                                                                                                                                                                                i.get("artist"),
                                                                                                                                                                                                i.get("movie_length"),
                                                                                                                                                                                                i.get("tags"),
                                                                                                                                                                                                int(i.get("update_time")),
                                                                                                                                                                                                int(i.get("create_time")))
            print sql1
            cursor.execute(sql1)
            db.commit()
        except Exception,e:
            db.rollback()
            print(e)
            
    for j in l:
        try:
            #insert
            sql1 = "UPDATE music_data SET id = %d ,movie_url = '%s' ,thumbnail_url = '%s' ,title = '%s' ,artist ='%s', movie_length = '%s', tags = '%s' ,update_time = %d ,create_time = %d" % (int(i.get("id")),
                                                                                                                                                                                                i.get("movie_url"),
                                                                                                                                                                                                i.get("thumbnail_url"),
                                                                                                                                                                                                i.get("title"),
                                                                                                                                                                                                i.get("artist"),
                                                                                                                                                                                                i.get("movie_length"),
                                                                                                                                                                                                i.get("tags"),
                                                                                                                                                                                                int(i.get("update_time")),
                                                                                                                                                                                                int(i.get("create_time")))
            print sql1
            cursor.execute(sql1)
            db.commit()
        except Exception,e:
            db.rollback()
            print(e)        
    db.close()
            
    print "===========finish"
    
    #print r.text    
if __name__=="__main__":
    #download_noteData()
    postSql_getData()
    
    
    #查看图片以及视频信息的api for nicovideo
    #http://ext.nicovideo.jp/api/getthumbinfo/sm