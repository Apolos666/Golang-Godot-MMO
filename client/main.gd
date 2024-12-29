extends Node

const packets := preload("res://packets.gd")

func _ready() -> void:
	var packet := packets.Packet.new()
	packet.set_sender_id(69)
	packet.new_chat().set_msg("Hello world")
	
	print(packet)
