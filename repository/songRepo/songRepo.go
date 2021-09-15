package songRepo

import (
	"context"
	"log"
	"songchord-api/driver"
	"songchord-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetSongByName(ctx context.Context, title string) (result interface{}) {
	var song models.Song
	data := driver.Mongo.ConnectCollection("song_chords", "songs").FindOne(ctx, bson.M{"title": title})
	data.Decode(&song)
	return song
}
func GetSongList(ctx context.Context, limit int) (result interface{}) {
	var song models.Song
	var songs []models.Song
	option := options.Find().SetLimit(int64(limit))
	cur, err := driver.Mongo.ConnectCollection("song_chords", "songs").Find(ctx, bson.M{}, option)
	defer cur.Close(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	for cur.Next(ctx) {
		cur.Decode(&song)
		songs = append(songs, song)
	}
	return songs
}
func InsertSong(ctx context.Context, song models.Song) error {
	_, err := driver.Mongo.ConnectCollection("song_chords", "songs").InsertOne(ctx, song)
	return err
}
func UpdateSong(ctx context.Context, song models.Song) error {
	filter := bson.M{"title": song.Title}
	update := bson.M{"$set": song}
	upsertBool := true
	updateOption := options.UpdateOptions{
		Upsert: &upsertBool,
	}
	_, err := driver.Mongo.ConnectCollection("song_chords", "songs").UpdateOne(ctx, filter, update, &updateOption)
	return err
}
func DeleteSong(ctx context.Context, title string) error {
	_, err := driver.Mongo.ConnectCollection("song_chords", "songs").DeleteOne(ctx, bson.M{"title": title})
	return err
}
